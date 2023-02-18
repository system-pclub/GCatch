
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[istio#8214]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[istio#8214]:(istio8214_test.go)
[patch]:https://github.com/istio/istio/pull/8214/files
[pull request]:https://github.com/istio/istio/pull/8214
 

## Backtrace

```
Write at 0x00c00011e088 by goroutine 10:
  sync/atomic.AddInt64()
      /usr/local/go/src/runtime/race_amd64.s:276 +0xb
  command-line-arguments.(*lruCache).SetWithExpiration()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:49 +0x43
  command-line-arguments.(*Cache).Set()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:24 +0x50
  command-line-arguments.(*grpcServer).check()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:58 +0xaa
  command-line-arguments.(*grpcServer).Check()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:63 +0x67
  command-line-arguments.TestIstio8214.func1.2()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:82 +0x66

Previous read at 0x00c00011e088 by goroutine 9:
  command-line-arguments.(*lruCache).Stats()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:41 +0x3a
  command-line-arguments.(*Cache).recordStats()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:29 +0x75
  command-line-arguments.(*Cache).Set()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:25 +0x51
  command-line-arguments.(*grpcServer).check()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:58 +0xaa
  command-line-arguments.(*grpcServer).Check()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:63 +0x67
  command-line-arguments.TestIstio8214.func1.1()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:78 +0x66

Goroutine 10 (running) created at:
  command-line-arguments.TestIstio8214.func1()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:80 +0x191

Goroutine 9 (finished) created at:
  command-line-arguments.TestIstio8214.func1()
      /root/gobench/goker/nonblocking/istio/8214/istio8214_test.go:76 +0x162
```

