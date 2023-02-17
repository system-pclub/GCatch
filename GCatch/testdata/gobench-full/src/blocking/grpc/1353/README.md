
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#1353]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[grpc#1353]:(grpc1353_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/1353/files
[pull request]:https://github.com/grpc/grpc-go/pull/1353
 
## Description

Some description from developers or pervious reseachers

> When it occurs?

> (1) roundRobin watchAddrUpdates sends an update to gRPC,
  and lbWatcher starts to process the update
> (2) roundRobin watchAddrUpdates sends another update to
  gRPC (while holding the mutex) this send blocks because the
  reader in lbWatcher is not reading. Also, the mutex is not
  released until the send unblocks.
>
> (3) lbWatcher calls down when processing the previous update.
  Since it removes some address, it tries to hold the mutex and blocks
  watchAddrUpdates is waiting for lbwatcher to read from the
  channel, while lbwatcher is waiting for watchAddrUpdates to
  release the mutex.

> The patch is to use an buffered channel and asks watchAddrUpdates
  to drain the channel before sending message, so that watchAddrUpdates
  will not be blocked at sending messages and it can release the lock.

Possible intervening

```
///
/// G1 					G2							G3
/// balancer.Start()
/// 					rr.watchAddrUpdates()
/// return
/// 												lbWatcher()
/// 												<-rr.addrCh
/// 					rr.mu.Lock()
/// 					rr.addrCh <- true
/// 					rr.mu.Unlock()
/// 												c.tearDown()
/// 												ac.down()
/// 					rr.mu.Lock()
/// 												rr.mu.Lock()
/// 					rr.addrCh <- true
/// ----------------------G2, G3 deadlock-----------------------
///
```

