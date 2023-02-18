
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#79631]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[kubernetes#79631]:(kubernetes79631_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/79631/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/79631
 

## Backtrace

```
Write at 0x00c00009c2d0 by goroutine 9:
  runtime.mapdelete_faststr()
      /usr/local/go/src/runtime/map_faststr.go:297 +0x0
  command-line-arguments.(*heapData).Pop()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:13 +0x6c
  command-line-arguments.Pop()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:21 +0xbc
  command-line-arguments.(*Heap).Pop()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:29 +0x98
  command-line-arguments.(*PriorityQueue).flushBackoffQCompleted()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:57 +0x78
  command-line-arguments.(*PriorityQueue).flushBackoffQCompleted-fm()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:54 +0x41
  command-line-arguments.BackoffUntil.func1()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:87 +0x57
  command-line-arguments.BackoffUntil()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:88 +0x4d
  command-line-arguments.JitterUntil()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:98 +0x43
  command-line-arguments.Until()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:102 +0x2b

Previous read at 0x00c00009c2d0 by goroutine 8:
  runtime.mapaccess1_faststr()
      /usr/local/go/src/runtime/map_faststr.go:12 +0x0
  command-line-arguments.(*Heap).GetByKey()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:37 +0xc8
  command-line-arguments.(*Heap).Get()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:33 +0x75
  command-line-arguments.TestKubernetes79631.func1()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:111 +0x56

Goroutine 9 (running) created at:
  command-line-arguments.(*PriorityQueue).run()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:75 +0xc3
  command-line-arguments.NewPriorityQueueWithClock()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:70 +0x17f
  command-line-arguments.NewPriorityQueue()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:62 +0x4c
  command-line-arguments.TestKubernetes79631.func1()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:110 +0x47

Goroutine 8 (finished) created at:
  command-line-arguments.TestKubernetes79631()
      /root/gobench/goker/nonblocking/kubernetes/79631/kubernetes79631_test.go:108 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

