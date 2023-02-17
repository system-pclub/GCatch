
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#9446]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[etcd#9446]:(etcd9446_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/9446/files
[pull request]:https://github.com/etcd-io/etcd/pull/9446
 

## Backtrace

```
  runtime.mapdelete_faststr()
      /usr/local/go/src/runtime/map_faststr.go:297 +0x0
  command-line-arguments.(*txBuffer).reset()
      /root/gobench/goker/nonblocking/etcd/9446/etcd9446_test.go:14 +0x106
  command-line-arguments.(*readTx).reset()
      /root/gobench/goker/nonblocking/etcd/9446/etcd9446_test.go:29 +0x6f
  command-line-arguments.TestEtcd9446.func1.1()
      /root/gobench/goker/nonblocking/etcd/9446/etcd9446_test.go:51 +0x66

Previous read at 0x00c00001c300 by goroutine 10:
  runtime.mapaccess1_faststr()
      /usr/local/go/src/runtime/map_faststr.go:12 +0x0
  command-line-arguments.(*txReadBuffer).Range()
      /root/gobench/goker/nonblocking/etcd/9446/etcd9446_test.go:21 +0xad
  command-line-arguments.(*readTx).UnsafeRange()
      /root/gobench/goker/nonblocking/etcd/9446/etcd9446_test.go:33 +0x65
  command-line-arguments.TestEtcd9446.func1.2()
      /root/gobench/goker/nonblocking/etcd/9446/etcd9446_test.go:55 +0x6f

Goroutine 9 (running) created at:
  command-line-arguments.TestEtcd9446.func1()
      /root/gobench/goker/nonblocking/etcd/9446/etcd9446_test.go:49 +0x138

Goroutine 10 (finished) created at:
  command-line-arguments.TestEtcd9446.func1()
      /root/gobench/goker/nonblocking/etcd/9446/etcd9446_test.go:53 +0x167
```

