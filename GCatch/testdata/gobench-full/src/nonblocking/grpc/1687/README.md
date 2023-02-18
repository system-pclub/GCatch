
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#1687]|[pull request]|[patch]| NonBlocking | Go-Specific | Misuse channel |

[grpc#1687]:(grpc1687_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/1687/files
[pull request]:https://github.com/grpc/grpc-go/pull/1687
 

## Backtrace

```
panic: send on closed channel

goroutine 7 [running]:
command-line-arguments.(*serverHandlerTransport).do(0xc00004c490, 0x60d788)
	/root/gobench/goker/nonblocking/grpc/1687/grpc1687_test.go:28 +0x15b
command-line-arguments.(*serverHandlerTransport).Write(0xc00004c490)
	/root/gobench/goker/nonblocking/grpc/1687/grpc1687_test.go:43 +0x45
command-line-arguments.TestGrpc1687.func1(0xc00004c4a0)
	/root/gobench/goker/nonblocking/grpc/1687/grpc1687_test.go:104 +0x76
created by command-line-arguments.testHandlerTransportHandleStreams.func1
	/root/gobench/goker/nonblocking/grpc/1687/grpc1687_test.go:98 +0x53
```

