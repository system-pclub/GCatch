
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#89164]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[kubernetes#89164]:(kubernetes89164_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/89164/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/89164
 

## Backtrace

```
Write at 0x00c00011e2e8 by goroutine 8:
  command-line-arguments.(*Cacher).startDispatching()
      /root/gobench/goker/nonblocking/kubernetes/89164/kubernetes89164_test.go:19 +0x93
  command-line-arguments.(*Cacher).dispatchEvent()
      /root/gobench/goker/nonblocking/kubernetes/89164/kubernetes89164_test.go:23 +0x38
  command-line-arguments.(*Cacher).dispatchEvents()
      /root/gobench/goker/nonblocking/kubernetes/89164/kubernetes89164_test.go:29 +0x38

Previous read at 0x00c00011e2e8 by goroutine 11:
  command-line-arguments.(*Cacher).dispatchEvent()
      /root/gobench/goker/nonblocking/kubernetes/89164/kubernetes89164_test.go:24 +0x4c
  command-line-arguments.TestKubernetes89164.func1()
      /root/gobench/goker/nonblocking/kubernetes/89164/kubernetes89164_test.go:48 +0x38

Goroutine 8 (running) created at:
  command-line-arguments.NewCacherFromConfig()
      /root/gobench/goker/nonblocking/kubernetes/89164/kubernetes89164_test.go:34 +0x93
  command-line-arguments.newTestCacher()
      /root/gobench/goker/nonblocking/kubernetes/89164/kubernetes89164_test.go:39 +0x34
  command-line-arguments.TestKubernetes89164()
      /root/gobench/goker/nonblocking/kubernetes/89164/kubernetes89164_test.go:43 +0x2f
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 11 (finished) created at:
  command-line-arguments.TestKubernetes89164()
      /root/gobench/goker/nonblocking/kubernetes/89164/kubernetes89164_test.go:47 +0xcb
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

