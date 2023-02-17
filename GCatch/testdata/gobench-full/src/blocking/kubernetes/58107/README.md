
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#58107]|[pull request]|[patch]| Blocking | Resource Deadlock | RWR deadlock |

[kubernetes#58107]:(kubernetes58107_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/58107/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/58107
 
## Description

Some description from developers or pervious reseachers

> The rules for read and write lock: allows concurrent read lock;
  write lock has higher priority than read lock.
  
> There are two queues (queue 1 and queue 2) involved in this bug,
  and the two queues are protected by the same read-write lock
  (rq.workerLock.RLock()). Before getting an element from queue 1 or
  queue 2, rq.workerLock.RLock() is acquired. If the queue is empty,
  cond.Wait() will be invoked. There is another goroutine (goroutine D),
  which will periodically invoke rq.workerLock.Lock(). Under the following
  situation, deadlock will happen. Queue 1 is empty, so that some goroutines
  hold rq.workerLock.RLock(), and block at cond.Wait(). Goroutine D is
  blocked when acquiring rq.workerLock.Lock(). Some goroutines try to process
  jobs in queue 2, but they are blocked when acquiring rq.workerLock.RLock(),
  since write lock has a higher priority.

> The fix is to not acquire rq.workerLock.RLock(), while pulling data
  from any queue. Therefore, when a goroutine is blocked at cond.Wait(),
  rq.workLock.RLock() is not held.

Possible intervening

```
/// G1 						G2						G3
/// ...						...						Sync()
/// rq.workerLock.RLock()
/// q.cond.Wait()
/// 												rq.workerLock.Lock()
/// 						rq.workerLock.RLock()
///
```

