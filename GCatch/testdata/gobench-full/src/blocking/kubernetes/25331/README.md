
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#25331]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel & Context |

[kubernetes#25331]:(kubernetes25331_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/25331/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/25331
 
## Description

Some description from developers or pervious reseachers

> In reflector.go, it could probably call Stop() without retrieving
  all results from ResultChan(). See here. A potential leak is that
  when an error has happened, it could block on resultChan, and then
  cancelling context in Stop() wouldn't unblock it.

Possible intervening

```
///
/// G1					G2
/// wc.run()
///						wc.Stop()
///						wc.errChan <-
///						wc.cancel()
///	<-wc.errChan
///	wc.cancel()
///	wc.resultChan <-
///	-------------G1 leak----------------
///

```

