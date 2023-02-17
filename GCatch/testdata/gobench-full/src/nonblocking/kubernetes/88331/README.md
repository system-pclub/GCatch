
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#88331]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[kubernetes#88331]:(kubernetes88331_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/88331/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/88331
 

## Backtrace

```
Write at 0x00c00000e0a0 by goroutine 9:
  command-line-arguments.(*data).Pop()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:13 +0x79
  command-line-arguments.Pop()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:21 +0xbc
  command-line-arguments.(*Heap).Pop()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:29 +0x98
  command-line-arguments.(*PriorityQueue).flushBackoffQCompleted()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:56 +0x78
  command-line-arguments.(*PriorityQueue).flushBackoffQCompleted-fm()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:53 +0x41
  command-line-arguments.BackoffUntil.func1()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:87 +0x57
  command-line-arguments.BackoffUntil()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:88 +0x4d
  command-line-arguments.JitterUntil()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:98 +0x43
  command-line-arguments.Until()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:102 +0x2b

Previous read at 0x00c00000e0a0 by goroutine 8:
  command-line-arguments.(*Heap).Len()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:32 +0x8f
  command-line-arguments.TestKubernetes88331.func1()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:111 +0x55

Goroutine 9 (running) created at:
  command-line-arguments.(*PriorityQueue).Run()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:75 +0xc3
  command-line-arguments.createAndRunPriorityQueue()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:70 +0x23e
  command-line-arguments.TestKubernetes88331.func1()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:110 +0x4b

Goroutine 8 (finished) created at:
  command-line-arguments.TestKubernetes88331()
      /root/gobench/goker/nonblocking/kubernetes/88331/kubernetes88331_test.go:108 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

