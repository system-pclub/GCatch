
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#81148]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[kubernetes#81148]:(kubernetes81148_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/81148/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/81148
 

## Backtrace

```
Read at 0x00c00001c370 by goroutine 9:
  command-line-arguments.(*PriorityQueue).flushUnschedulableQLeftover()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:50 +0x169
  command-line-arguments.(*PriorityQueue).flushUnschedulableQLeftover-fm()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:45 +0x41
  command-line-arguments.BackoffUntil.func1()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:87 +0x57
  command-line-arguments.BackoffUntil()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:88 +0x4d
  command-line-arguments.JitterUntil()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:98 +0x43
  command-line-arguments.Until()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:102 +0x2b

Previous write at 0x00c00001c370 by goroutine 8:
  command-line-arguments.TestKubernetes81148.func1()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:119 +0x1a0

Goroutine 9 (running) created at:
  command-line-arguments.(*PriorityQueue).run()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:55 +0xc3
  command-line-arguments.NewPriorityQueueWithClock()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:70 +0x149
  command-line-arguments.NewPriorityQueue()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:75 +0x87
  command-line-arguments.TestKubernetes81148.func1()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:116 +0x7a

Goroutine 8 (finished) created at:
  command-line-arguments.TestKubernetes81148()
      /root/gobench/goker/nonblocking/kubernetes/81148/kubernetes81148_test.go:114 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

