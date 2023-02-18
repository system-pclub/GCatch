
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#27782]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel & Condition Variable |

[moby#27782]:(moby27782_test.go)
[patch]:https://github.com/moby/moby/pull/27782/files
[pull request]:https://github.com/moby/moby/pull/27782
 
## Description

Possible intervening

```
///
/// G1 						G2							G3
/// InitializeStdio()
/// startLogging()
/// l.ReadLogs()
/// NewLogWatcher()
/// 						l.readLogs()
/// container.Reset()
/// LogDriver.Close()
/// r.Close()
/// close(w.closeNotifier)
/// 						followLogs(logWatcher)
/// 						watchFile()
/// 						New()
/// 						NewEventWatcher()
/// 						NewWatcher()
/// 													w.readEvents()
/// 													event.ignoreLinux()
/// 													return false
/// 						<-logWatcher.WatchClose()
/// 						fileWatcher.Remove()
/// 						w.cv.Wait()
/// 													w.Events <- event
/// ------------------------------G2,G3 deadlock---------------------------
///

```
