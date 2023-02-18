
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#2371]|[pull request]|[patch]| NonBlocking | Go-Specific | Misuse channel |

[grpc#2371]:(grpc2371_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/2371/files
[pull request]:https://github.com/grpc/grpc-go/pull/2371
 

## Backtrace

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x5b43b4]

goroutine 7 [running]:
command-line-arguments.(*ccBalancerWrapper).handleResolvedAddrs(0x0)
	/root/gobench/goker/nonblocking/grpc/2371/grpc2371_test.go:16 +0x34
command-line-arguments.(*ClientConn).handleServiceConfig(0xc00001c300)
	/root/gobench/goker/nonblocking/grpc/2371/grpc2371_test.go:57 +0x5f
command-line-arguments.(*ccResolverWrapper).watcher(0xc000010038)
	/root/gobench/goker/nonblocking/grpc/2371/grpc2371_test.go:39 +0x4c
created by command-line-arguments.(*ccResolverWrapper).start
	/root/gobench/goker/nonblocking/grpc/2371/grpc2371_test.go:35 +0x4d
```

