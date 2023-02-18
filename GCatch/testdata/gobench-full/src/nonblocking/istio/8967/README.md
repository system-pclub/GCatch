
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[istio#8967]|[pull request]|[patch]| NonBlocking | Go-Specific | Misuse channel |

[istio#8967]:(istio8967_test.go)
[patch]:https://github.com/istio/istio/pull/8967/files
[pull request]:https://github.com/istio/istio/pull/8967
 

## Backtrace

```
Read at 0x00c0000aa028 by goroutine 9:
  command-line-arguments.(*fsSource).Start.func1()
      /root/gobench/goker/nonblocking/istio/8967/istio8967_test.go:22 +0x3a

Previous write at 0x00c0000aa028 by goroutine 8:
  command-line-arguments.(*fsSource).Stop()
      /root/gobench/goker/nonblocking/istio/8967/istio8967_test.go:31 +0x59
  command-line-arguments.TestIstio8967.func1()
      /root/gobench/goker/nonblocking/istio/8967/istio8967_test.go:51 +0xe6

Goroutine 9 (running) created at:
  command-line-arguments.(*fsSource).Start()
      /root/gobench/goker/nonblocking/istio/8967/istio8967_test.go:19 +0x4c
  command-line-arguments.TestIstio8967.func1()
      /root/gobench/goker/nonblocking/istio/8967/istio8967_test.go:50 +0xd8

Goroutine 8 (running) created at:
  command-line-arguments.TestIstio8967()
      /root/gobench/goker/nonblocking/istio/8967/istio8967_test.go:47 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

