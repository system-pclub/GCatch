
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#6873]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[etcd#6873]:(etcd6873_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/6873/files
[pull request]:https://github.com/etcd-io/etcd/pull/6873
 
## Description

Possible intervening

```
///
/// G1						G2					G3
/// newWatchBroadcasts()
///	wbs.update()
/// wbs.updatec <-
/// return
///							<-wbs.updatec
///							wbs.coalesce()
///												wbs.stop()
///												wbs.mu.Lock()
///												close(wbs.updatec)
///												<-wbs.donec
///							wbs.mu.Lock()
///---------------------G2,G3 deadlock-------------------------
///
```

