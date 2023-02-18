# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#1462]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & WaitGroup |

[cockroach#1462]:(cockroach1462_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/1462/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/1462
 
## Description

`s.stop.Wait()` in `Stop()` not guaranteed to be invoked

### backtrace

```
goroutine 12 [chan receive]:
goroutine 4613 [semacquire, 9 minutes]:
sync.runtime_Semacquire(0xc0000ea418)
	/usr/local/go/src/runtime/sema.go:56 +0x42
sync.(*WaitGroup).Wait(0xc0000ea410)
	/usr/local/go/src/sync/waitgroup.go:130 +0x64
command-line-arguments.(*Stopper).Stop(0xc0000ea400)
	/root/gobench/gobench/goker/blocking/cockroach/1462/cockroach1462_test.go:79 +0x5f
command-line-arguments.TestCockroach1462(0xc0000e7320)
	/root/gobench/gobench/goker/blocking/cockroach/1462/cockroach1462_test.go:139 +0x1fc
testing.tRunner(0xc0000e7320, 0x554420)
	/usr/local/go/src/testing/testing.go:1050 +0xdc
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1095 +0x28b

goroutine 4614 [chan send, 9 minutes]:
command-line-arguments.(*localInterceptableTransport).start.func1()
	/root/gobench/gobench/goker/blocking/cockroach/1462/cockroach1462_test.go:115 +0x46
command-line-arguments.(*Stopper).RunWorker.func1(0xc0000ea400, 0xc0000e42c0)
	/root/gobench/gobench/goker/blocking/cockroach/1462/cockroach1462_test.go:31 +0x4f
created by command-line-arguments.(*Stopper).RunWorker
	/root/gobench/gobench/goker/blocking/cockroach/1462/cockroach1462_test.go:29 +0x67
```