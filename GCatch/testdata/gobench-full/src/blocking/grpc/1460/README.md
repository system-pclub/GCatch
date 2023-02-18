
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#1460]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[grpc#1460]:(grpc1460_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/1460/files
[pull request]:https://github.com/grpc/grpc-go/pull/1460
 
## Description

Some description from developers or pervious reseachers

> When gRPC keepalives are enabled (which isn't the case
  by default at this time) and PermitWithoutStream is false
  (the default), the client can deadlock when transitioning
  between having no active stream and having one active
  stream.The keepalive() goroutine is stuck at “<-t.awakenKeepalive”,
  while the main goroutine is stuck in NewStream() on t.mu.Lock().

Possible intervening

```
///
/// G1 						G2
/// client.keepalive()
/// 						client.NewStream()
/// t.mu.Lock()
/// <-t.awakenKeepalive
/// 						t.mu.Lock()
/// ---------------G1, G2 deadlock--------------
///
```

