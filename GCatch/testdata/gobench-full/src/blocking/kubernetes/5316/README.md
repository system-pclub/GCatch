
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#5316]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[kubernetes#5316]:(kubernetes5316_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/5316/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/5316
 
## Description

Some description from developers or pervious reseachers

> If the main goroutine selects a case that doesn’t consumes
  the channels, the anonymous goroutine will be blocked on sending
  to channel.

Possible intervening

```
///
/// G1 						G2
/// finishRequest()
/// 						fn()
/// time.After()
/// 						errCh<-/ch<-
/// --------------G2 leak----------------
///
```

