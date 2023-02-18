
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#81091]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[kubernetes#81091]:(kubernetes81091_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/81091/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/81091
 

## Backtrace

```
Read at 0x00c000186ce8 by goroutine 26:
  command-line-arguments.(*FakeFilterPlugin).Filter()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:13 +0x3a
  command-line-arguments.(*framework).RunFilterPlugins()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:36 +0x90
  command-line-arguments.(*genericScheduler).findNodesThatFit.func1()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:52 +0x5e
  command-line-arguments.ParallelizeUntil.func1()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:86 +0x80

Previous write at 0x00c000186ce8 by goroutine 10:
  command-line-arguments.(*FakeFilterPlugin).Filter()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:13 +0x50
  command-line-arguments.(*framework).RunFilterPlugins()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:36 +0x90
  command-line-arguments.(*genericScheduler).findNodesThatFit.func1()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:52 +0x5e
  command-line-arguments.ParallelizeUntil.func1()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:86 +0x80

Goroutine 26 (running) created at:
  command-line-arguments.ParallelizeUntil()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:79 +0x189
  command-line-arguments.(*genericScheduler).findNodesThatFit()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:54 +0xa4
  command-line-arguments.(*genericScheduler).Schedule()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:58 +0x22c
  command-line-arguments.TestKubernetes81091.func1()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:101 +0x223

Goroutine 10 (finished) created at:
  command-line-arguments.ParallelizeUntil()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:79 +0x189
  command-line-arguments.(*genericScheduler).findNodesThatFit()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:54 +0xa4
  command-line-arguments.(*genericScheduler).Schedule()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:58 +0x22c
  command-line-arguments.TestKubernetes81091.func1()
      /root/gobench/goker/nonblocking/kubernetes/81091/kubernetes81091_test.go:101 +0x223
```

