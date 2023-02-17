
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#862]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel & Context |

[grpc#862]:(grpc862_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/862/files
[pull request]:https://github.com/grpc/grpc-go/pull/862
 
## Description

Some description from developers or pervious reseachers

> When return value conn is nil, cc (ClientConn) is not closed.
  The goroutine executing resetAddrConn is leaked. The patch is to
  close ClientConn in the defer func().

Possible intervening

```
///
/// G1 					G2
/// DialContext()
/// 					cc.resetAddrConn()
/// 					resetTransport()
/// 					<-ac.ctx.Done()
/// --------------G2 leak------------------
///
```

