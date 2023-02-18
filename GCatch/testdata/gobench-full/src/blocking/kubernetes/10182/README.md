
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#10182]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[kubernetes#10182]:(kubernetes10182_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/10182/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/10182
 
## Description

Some description from developers or pervious reseachers

>  This is a lock-channel bug. goroutine 1 is blocked on a lock
   held by goroutine 3, while goroutine 3 is blocked on sending
   message to ch, which is read by goroutine 1.

Possible intervening

```
/// G1 						G2							G3
/// s.Start()
/// s.syncBatch()
/// 						s.SetPodStatus()
/// <-s.podStatusChannel
/// 						s.podStatusesLock.Lock()
/// 						s.podStatusChannel <- true
/// 						s.podStatusesLock.Unlock()
/// 						return
/// s.DeletePodStatus()
/// 													s.podStatusesLock.Lock()
/// 													s.podStatusChannel <- true
/// s.podStatusesLock.Lock()
/// -----------------------------G1,G3 deadlock----------------------------
```

