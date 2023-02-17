
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#660]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[grpc#660]:(grpc660_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/660/files
[pull request]:https://github.com/grpc/grpc-go/pull/660
 
## Description

Some description from developers or pervious reseachers

> The parent function could return without draining the done channel.

Possible intervening

```
///
/// G1 						G2 				helper goroutine
/// doCloseLoopUnary()
///											bc.stop <- true
/// <-bc.stop
/// return
/// 						done <-
/// ----------------------G2 leak--------------------------
///
```

