
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#8194]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[etcd#8194]:(etcd8194_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/8194/files
[pull request]:https://github.com/etcd-io/etcd/pull/8194
 

## Backtrace

```
Write at 0x000000737418 by goroutine 9:
  command-line-arguments.testLessorRenewExtendPileup()
      /root/gobench/goker/nonblocking/etcd/8194/etcd8194_test.go:14 +0x80
  command-line-arguments.TestEtcd8194.func2()
      /root/gobench/goker/nonblocking/etcd/8194/etcd8194_test.go:72 +0x5f

Previous read at 0x000000737418 by goroutine 10:
  command-line-arguments.(*lessor).runLoop()
      /root/gobench/goker/nonblocking/etcd/8194/etcd8194_test.go:35 +0x282

Goroutine 9 (running) created at:
  command-line-arguments.TestEtcd8194()
      /root/gobench/goker/nonblocking/etcd/8194/etcd8194_test.go:70 +0xc4
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 10 (running) created at:
  command-line-arguments.newLessor()
      /root/gobench/goker/nonblocking/etcd/8194/etcd8194_test.go:55 +0x93
  command-line-arguments.testLessorGrant()
      /root/gobench/goker/nonblocking/etcd/8194/etcd8194_test.go:60 +0x60
  command-line-arguments.TestEtcd8194.func1()
      /root/gobench/goker/nonblocking/etcd/8194/etcd8194_test.go:68 +0x5b
```

