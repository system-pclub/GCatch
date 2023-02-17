
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[serving#2137]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[serving#2137]:(serving2137_test.go)
[patch]:https://github.com/ knative/serving/pull/2137/files
[pull request]:https://github.com/ knative/serving/pull/2137
 
## Description

Possible intervening

```
//
// G1                           G2                      G3
// b.concurrentRequests(2)
// b.concurrentRequest()
// r.lock.Lock()
//                                                      start.Done()
// start.Wait()
// b.concurrentRequest()
// r.lock.Lock()
//                              start.Done()
// start.Wait()
// unlockAll(locks)
// unlock(lc)
// req.lock.Unlock()
// ok := <-req.accepted
//                              b.Maybe()
//                              b.activeRequests <- t
//                              thunk()
//                              r.lock.Lock()
//                                                      b.Maybe()
//                                                      b.activeRequests <- t
// ----------------------------G1,G2,G3 deadlock-----------------------------
//
```

