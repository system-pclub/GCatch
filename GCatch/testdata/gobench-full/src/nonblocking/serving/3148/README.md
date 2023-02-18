
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[serving#3148]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[serving#3148]:(serving3148_test.go)
[patch]:https://github.com/ knative/serving/pull/3148/files
[pull request]:https://github.com/ knative/serving/pull/3148
 

## Backtrace

```
Read at 0x00c0000b6040 by goroutine 10:
  command-line-arguments.(*Fake).Invokes()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:148 +0x3a
  command-line-arguments.(*FakePodAutoscalers).Create()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:45 +0x63
  command-line-arguments.(*Reconciler).createKPA()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:74 +0x78
  command-line-arguments.(*Reconciler).reconcileKPA()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:70 +0x41
  command-line-arguments.(*Reconciler).reconcileKPA-fm()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:69 +0x22
  command-line-arguments.(*Reconciler).reconcile()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:65 +0xbb
  command-line-arguments.(*Reconciler).Reconcile()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:53 +0x38
  command-line-arguments.(*Impl).processNextWorkItem()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:99 +0x85
  command-line-arguments.(*Impl).Run.func1()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:93 +0x66

Previous write at 0x00c0000b6040 by goroutine 8:
  command-line-arguments.(*Fake).PrependReactor()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:153 +0x31f
  command-line-arguments.(*Hooks).OnUpdate()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:136 +0x2c2
  command-line-arguments.TestServing3148.func1()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:170 +0x2c1

Goroutine 10 (running) created at:
  command-line-arguments.(*Impl).Run()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:91 +0xfd
  command-line-arguments.TestServing3148.func1.2()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:168 +0x4a
  command-line-arguments.(*Group).Go.func1()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:126 +0x6a

Goroutine 8 (running) created at:
  command-line-arguments.TestServing3148()
      /root/gobench/goker/nonblocking/serving/3148/serving3148_test.go:159 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

