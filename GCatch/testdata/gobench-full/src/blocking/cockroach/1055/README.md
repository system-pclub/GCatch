# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#1055]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & WaitGroup |

[cockroach#1055]:(cockroach1055_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/1055/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/1055
 
## Description


This is some description from developers

> 1. Stop() is called and blocked at s.stop.Wait() after acquiring the lock.
> 2. StartTask() is called and attempts to acquire the lock. It is then blocked.
> 3. Stop() never finishes since the task doesn't call SetStopped.


### backtrace

```
goroutine 16 [semacquire]:
sync.runtime_Semacquire(0xc00001a778)
	/usr/local/go/src/runtime/sema.go:56 +0x42
sync.(*WaitGroup).Wait(0xc00001a770)
	/usr/local/go/src/sync/waitgroup.go:130 +0x64
command-line-arguments.(*Stopper).Stop(0xc00001a750)
	/root/gobench/gobench/goker/blocking/cockroach/1055/cockroach1055_test.go:46 +0x70
command-line-arguments.TestCockroach1055.func2(0xc00000c0c0, 0x3, 0x4, 0xc0000628a0)
	/root/gobench/gobench/goker/blocking/cockroach/1055/cockroach1055_test.go:89 +0x69
created by command-line-arguments.TestCockroach1055
	/root/gobench/gobench/goker/blocking/cockroach/1055/cockroach1055_test.go:84 +0x29c

goroutine 15 [chan receive]:
command-line-arguments.TestCockroach1055.func1(0xc00001a750)
	/root/gobench/gobench/goker/blocking/cockroach/1055/cockroach1055_test.go:78 +0x4a
created by command-line-arguments.TestCockroach1055
	/root/gobench/gobench/goker/blocking/cockroach/1055/cockroach1055_test.go:76 +0x21c
```