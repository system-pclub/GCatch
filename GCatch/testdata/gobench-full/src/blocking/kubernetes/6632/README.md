
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#6632]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[kubernetes#6632]:(kubernetes6632_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/6632/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/6632
 
## Description

Some description from developers or pervious reseachers

> This is a lock-channel bug. When resetChan is full, WriteFrame
  holds the lock and blocks on the channel. Then monitor() fails
  to close the resetChan because lock is already held by WriteFrame.
  
> Fix: create a goroutine to drain the channel

Possible intervening

```
///
/// G1						G2					helper goroutine
/// i.monitor()
/// <-i.conn.closeChan
///							i.WriteFrame()
///							i.writeLock.Lock()
///							i.resetChan <-
///												i.conn.closeChan<-
///	i.writeLock.Lock()
///	----------------------G1,G2 deadlock------------------------
///
```

