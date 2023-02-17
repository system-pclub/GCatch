
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#4876]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[etcd#4876]:(etcd4876_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/4876/files
[pull request]:https://github.com/etcd-io/etcd/pull/4876
 

## Backtrace

```
Read at 0x000000734418 by goroutine 11:
  command-line-arguments.(*serverWatchStream).sendLoop()
      /root/gobench/goker/nonblocking/etcd/4876/etcd4876_test.go:33 +0x3a

Previous write at 0x000000734418 by goroutine 9:
  command-line-arguments.TestEtcd4876.func1.1()
      /root/gobench/goker/nonblocking/etcd/4876/etcd4876_test.go:52 +0x6e

Goroutine 11 (running) created at:
  command-line-arguments.(*watchServer).Watch()
      /root/gobench/goker/nonblocking/etcd/4876/etcd4876_test.go:40 +0x4e
  command-line-arguments.TestEtcd4876.func1.2()
      /root/gobench/goker/nonblocking/etcd/4876/etcd4876_test.go:56 +0x98

Goroutine 9 (finished) created at:
  command-line-arguments.TestEtcd4876.func1()
      /root/gobench/goker/nonblocking/etcd/4876/etcd4876_test.go:49 +0x86
```

