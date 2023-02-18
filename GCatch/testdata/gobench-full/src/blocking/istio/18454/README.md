
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[istio#18454]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel & Context |

[istio#18454]:(istio18454_test.go)
[patch]:https://github.com/istio/istio/pull/18454/files
[pull request]:https://github.com/istio/istio/pull/18454
 
## Description

s.timer.Stop() at line 56 and 61 can be called concurrency 
(i.e. from their entry point at line 104 and line 66).
See [Timer](https://golang.org/pkg/time/#Timer).


## Backtrace

```
goroutine 5 [chan receive]:
command-line-arguments.(*Strategy).startTimer.func1(0x578320, 0xc000090800)
    /root/gobench/goker/blocking/istio/18454/istio18454_test.go:58 +0x17b
command-line-arguments.(*Worker).Start.func1(0xc00005e190, 0xc0000a4240)
    /root/gobench/goker/blocking/istio/18454/istio18454_test.go:22 +0x3c
created by command-line-arguments.(*Worker).Start
    /root/gobench/goker/blocking/istio/18454/istio18454_test.go:21 +0x53

 Goroutine 31 in state chan send, with command-line-arguments.(*Strategy).OnChange on top of the stack:
goroutine 31 [chan send]:
command-line-arguments.(*Strategy).OnChange(0xc0000970b0)
    /root/gobench/goker/blocking/istio/18454/istio18454_test.go:43 +0x74
command-line-arguments.(*Processor).processEvent(...)
    /root/gobench/goker/blocking/istio/18454/istio18454_test.go:83
command-line-arguments.(*Processor).Start.func2(0x578320, 0xc000090840)
    /root/gobench/goker/blocking/istio/18454/istio18454_test.go:101 +0x78
command-line-arguments.(*Worker).Start.func1(0xc000092530, 0xc0000a4260)
    /root/gobench/goker/blocking/istio/18454/istio18454_test.go:22 +0x3c
created by command-line-arguments.(*Worker).Start
    /root/gobench/goker/blocking/istio/18454/istio18454_test.go:21 +0x53
```

