
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#3077]|[pull request]|[patch]| NonBlocking | Go-Specific | Misuse channel |

[etcd#3077]:(etcd3077_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/3077/files
[pull request]:https://github.com/etcd-io/etcd/pull/3077
 

## Backtrace

```
Read at 0x00c00011e2d8 by goroutine 8:
  command-line-arguments.(*EtcdServer).run.func1()
      /root/gobench/goker/nonblocking/etcd/3077/etcd3077_test.go:41 +0x42
  command-line-arguments.(*EtcdServer).run()
      /root/gobench/goker/nonblocking/etcd/3077/etcd3077_test.go:49 +0xc2

Previous write at 0x00c00011e2d8 by goroutine 9:
  command-line-arguments.(*raftNode).run()
      /root/gobench/goker/nonblocking/etcd/3077/etcd3077_test.go:15 +0x81

Goroutine 8 (running) created at:
  command-line-arguments.(*EtcdServer).start()
      /root/gobench/goker/nonblocking/etcd/3077/etcd3077_test.go:57 +0xf3
  command-line-arguments.TestEtcd3077()
      /root/gobench/goker/nonblocking/etcd/3077/etcd3077_test.go:73 +0x83
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 9 (running) created at:
  command-line-arguments.(*EtcdServer).run()
      /root/gobench/goker/nonblocking/etcd/3077/etcd3077_test.go:37 +0x52
```

