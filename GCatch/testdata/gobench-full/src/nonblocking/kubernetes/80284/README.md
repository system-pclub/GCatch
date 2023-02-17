
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#80284]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[kubernetes#80284]:(kubernetes80284_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/80284/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/80284
 

## Backtrace

```
Write at 0x00c000010038 by goroutine 9:
  command-line-arguments.(*Authenticator).UpdateTransportConfig()
      /root/gobench/goker/nonblocking/kubernetes/80284/kubernetes80284_test.go:22 +0x9c
  command-line-arguments.TestKubernetes80284.func1()
      /root/gobench/goker/nonblocking/kubernetes/80284/kubernetes80284_test.go:36 +0x6c

Previous write at 0x00c000010038 by goroutine 8:
  command-line-arguments.(*Authenticator).UpdateTransportConfig()
      /root/gobench/goker/nonblocking/kubernetes/80284/kubernetes80284_test.go:22 +0x9c
  command-line-arguments.TestKubernetes80284.func1()
      /root/gobench/goker/nonblocking/kubernetes/80284/kubernetes80284_test.go:36 +0x6c

Goroutine 9 (running) created at:
  command-line-arguments.TestKubernetes80284()
      /root/gobench/goker/nonblocking/kubernetes/80284/kubernetes80284_test.go:34 +0xe9
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 8 (finished) created at:
  command-line-arguments.TestKubernetes80284()
      /root/gobench/goker/nonblocking/kubernetes/80284/kubernetes80284_test.go:34 +0xe9
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

