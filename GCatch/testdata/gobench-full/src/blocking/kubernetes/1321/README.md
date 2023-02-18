
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#1321]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[kubernetes#1321]:(kubernetes1321_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/1321/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/1321
 
## Description

Some description from developers or pervious reseachers

> This is a lock-channel bug. The first goroutine invokes
> distribute() function. distribute() function holds m.lock.Lock(),
  while blocking at sending message to w.result. The second goroutine
  invokes stopWatching() funciton, which can unblock the first
  goroutine by closing w.result. However, in order to close w.result,
  stopWatching() function needs to acquire m.lock.Lock() firstly.
>
> The fix is to introduce another channel and put receive message
  from the second channel in the same select as the w.result. Close
  the second channel can unblock the first goroutine, while no need
  to hold m.lock.Lock().

Possible intervening

```
///
/// G1 							G2
/// testMuxWatcherClose()
/// NewMux()
/// 							m.loop()
/// 							m.distribute()
/// 							m.lock.Lock()
/// 							w.result <- true
/// w := m.Watch()
/// w.Stop()
/// mw.m.stopWatching()
/// m.lock.Lock()
/// ---------------G1,G2 deadlock---------------
///
```

