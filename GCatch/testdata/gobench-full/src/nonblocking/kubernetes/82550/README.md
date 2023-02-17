
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#82550]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[kubernetes#82550]:(kubernetes82550_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/82550/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/82550
 

## Backtrace

```
Read at 0x00c0000aa028 by goroutine 9:
  command-line-arguments.(*lazyEcrProvider).LazyProvide()
      /root/gobench/goker/nonblocking/kubernetes/82550/kubernetes82550_test.go:24 +0x52

Previous write at 0x00c0000aa028 by goroutine 8:
  command-line-arguments.(*lazyEcrProvider).LazyProvide()
      /root/gobench/goker/nonblocking/kubernetes/82550/kubernetes82550_test.go:25 +0x14d

Goroutine 9 (running) created at:
  command-line-arguments.TestKubernetes82550()
      /root/gobench/goker/nonblocking/kubernetes/82550/kubernetes82550_test.go:34 +0x8c
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 8 (finished) created at:
  command-line-arguments.TestKubernetes82550()
      /root/gobench/goker/nonblocking/kubernetes/82550/kubernetes82550_test.go:34 +0x8c
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

```
Read at 0x00c0000aa028 by goroutine 10:
  command-line-arguments.(*lazyEcrProvider).LazyProvide()
      /root/gobench/goker/nonblocking/kubernetes/82550/kubernetes82550_test.go:24 +0x52

Previous write at 0x00c0000aa028 by goroutine 8:
  command-line-arguments.(*lazyEcrProvider).LazyProvide()
      /root/gobench/goker/nonblocking/kubernetes/82550/kubernetes82550_test.go:25 +0x14d

Goroutine 10 (running) created at:
  command-line-arguments.TestKubernetes82550()
      /root/gobench/goker/nonblocking/kubernetes/82550/kubernetes82550_test.go:34 +0x8c
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 8 (finished) created at:
  command-line-arguments.TestKubernetes82550()
      /root/gobench/goker/nonblocking/kubernetes/82550/kubernetes82550_test.go:34 +0x8c
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```
