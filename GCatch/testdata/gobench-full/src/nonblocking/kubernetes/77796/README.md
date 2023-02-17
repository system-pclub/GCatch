
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#77796]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[kubernetes#77796]:(kubernetes77796_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/77796/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/77796
 

## Backtrace

```
Write at 0x00c0001142e8 by goroutine 9:
  command-line-arguments.(*Cacher).startDispatching()
      /root/gobench/goker/nonblocking/kubernetes/77796/kubernetes77796_test.go:20 +0x93
  command-line-arguments.(*Cacher).dispatchEvent()
      /root/gobench/goker/nonblocking/kubernetes/77796/kubernetes77796_test.go:24 +0x38
  command-line-arguments.TestKubernetes77796.func1()
      /root/gobench/goker/nonblocking/kubernetes/77796/kubernetes77796_test.go:47 +0x38

Previous read at 0x00c0001142e8 by goroutine 8:
  command-line-arguments.(*Cacher).dispatchEvent()
      /root/gobench/goker/nonblocking/kubernetes/77796/kubernetes77796_test.go:25 +0x4c
  command-line-arguments.(*Cacher).dispatchEvents()
      /root/gobench/goker/nonblocking/kubernetes/77796/kubernetes77796_test.go:30 +0x38

Goroutine 9 (running) created at:
  command-line-arguments.TestKubernetes77796()
      /root/gobench/goker/nonblocking/kubernetes/77796/kubernetes77796_test.go:46 +0x63
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 8 (finished) created at:
  command-line-arguments.NewCacherFromConfig()
      /root/gobench/goker/nonblocking/kubernetes/77796/kubernetes77796_test.go:35 +0x93
  command-line-arguments.newTestCacher()
      /root/gobench/goker/nonblocking/kubernetes/77796/kubernetes77796_test.go:40 +0x34
  command-line-arguments.TestKubernetes77796()
      /root/gobench/goker/nonblocking/kubernetes/77796/kubernetes77796_test.go:44 +0x2f
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

