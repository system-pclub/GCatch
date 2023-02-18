
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#1424]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[grpc#1424]:(grpc1424_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/1424/files
[pull request]:https://github.com/grpc/grpc-go/pull/1424
 
## Description

Some description from developers or pervious reseachers

> The parent function could return without draining the done channel.

Possible intervening

```
///
/// G1                      G2                          G3
/// DialContext()
///                         cc.dopts.balancer.Notify()
///                                                     cc.lbWatcher()
///                         <-doneChan
/// close()
/// -----------------------G2 leak------------------------------------
///
```

