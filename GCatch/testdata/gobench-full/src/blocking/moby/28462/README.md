
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#28462]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[moby#28462]:(moby28462_test.go)
[patch]:https://github.com/moby/moby/pull/28462/files
[pull request]:https://github.com/moby/moby/pull/28462
 
## Description

Some description from developers or pervious reseachers

> There are three goroutines mentioned in the bug report Moby#28405.
  Actually, only two goroutines are needed to trigger this bug. This bug
  is another example where lock and channel are mixed with each other.

Possible intervening

```
///
/// G1							G2
/// monitor()
/// handleProbeResult()
/// 							d.StateChanged()
/// 							c.Lock()
/// 							d.updateHealthMonitorElseBranch()
/// 							h.CloseMonitorChannel()
/// 							s.stop <- struct{}{}
/// c.Lock()
/// ----------------------G1,G2 deadlock------------------------
///
```

