
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#3090]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[grpc#3090]:(grpc3090_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/3090/files
[pull request]:https://github.com/grpc/grpc-go/pull/3090
 

## Backtrace

```
Read at 0x00c0000c0048 by goroutine 9:
  command-line-arguments.(*ccResolverWrapper).resolveNow()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:42 +0x55
  command-line-arguments.(*ccResolverWrapper).poll.func1()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:50 +0x38

Previous write at 0x00c0000c0048 by goroutine 8:
  command-line-arguments.newCCResolverWrapper()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:61 +0x172
  command-line-arguments.DialContext()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:83 +0x9f
  command-line-arguments.Dial()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:86 +0x64
  command-line-arguments.TestGrpc3090.func1()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:94 +0x5f

Goroutine 9 (running) created at:
  command-line-arguments.(*ccResolverWrapper).poll()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:49 +0x9f
  command-line-arguments.(*ccResolverWrapper).UpdateState()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:55 +0x38
  command-line-arguments.(*resolver_Resolver).UpdateState()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:27 +0x7d
  command-line-arguments.(*resolver_Resolver).Build()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:19 +0x5e
  command-line-arguments.newCCResolverWrapper()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:61 +0x14c
  command-line-arguments.DialContext()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:83 +0x9f
  command-line-arguments.Dial()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:86 +0x64
  command-line-arguments.TestGrpc3090.func1()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:94 +0x5f

Goroutine 8 (running) created at:
  command-line-arguments.TestGrpc3090()
      /root/gobench/goker/nonblocking/grpc/3090/grpc3090_test.go:92 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

