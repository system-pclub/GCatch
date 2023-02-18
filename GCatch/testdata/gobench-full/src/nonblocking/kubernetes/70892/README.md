
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#70892]|[pull request]|[patch]| NonBlocking | Go-Specific | Anonymous function |

[kubernetes#70892]:(kubernetes70892_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/70892/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/70892
 

## Backtrace

```
Write at 0x00c00006a870 by goroutine 64:
  command-line-arguments.TestKubernetes70892.func1()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:57 +0x190
  command-line-arguments.ParallelizeUntil.func1()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:39 +0x80

Previous read at 0x00c00006a870 by goroutine 61:
  command-line-arguments.TestKubernetes70892.func1()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:56 +0xb3
  command-line-arguments.ParallelizeUntil.func1()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:39 +0x80

Goroutine 64 (running) created at:
  command-line-arguments.ParallelizeUntil()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:32 +0x1a6
  command-line-arguments.TestKubernetes70892()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:61 +0x2e5
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 61 (running) created at:
  command-line-arguments.ParallelizeUntil()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:32 +0x1a6
  command-line-arguments.TestKubernetes70892()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:61 +0x2e5
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

```
Read at 0x00c000091640 by goroutine 61:
  command-line-arguments.TestKubernetes70892.func1()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:56 +0xdd
  command-line-arguments.ParallelizeUntil.func1()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:39 +0x80

Previous write at 0x00c000091640 by goroutine 64:
  command-line-arguments.TestKubernetes70892.func1()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:57 +0x10f
  command-line-arguments.ParallelizeUntil.func1()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:39 +0x80

Goroutine 61 (running) created at:
  command-line-arguments.ParallelizeUntil()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:32 +0x1a6
  command-line-arguments.TestKubernetes70892()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:61 +0x2e5
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 64 (finished) created at:
  command-line-arguments.ParallelizeUntil()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:32 +0x1a6
  command-line-arguments.TestKubernetes70892()
      /root/gobench/goker/nonblocking/kubernetes/70892/kubernetes70892_test.go:61 +0x2e5
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```
