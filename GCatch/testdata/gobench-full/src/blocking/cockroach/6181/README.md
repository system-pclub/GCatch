# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#6181]|[pull request]|[patch]| Blocking | Resource Deadlock | RWR Deadlock |

[cockroach#6181]:(cockroach6181_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/6181/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/6181
 
## Description

Possible intervening

```
/// G1 									G2							G3					...
/// testRangeCacheCoalescedRquests()
/// initTestDescriptorDB()
/// pauseLookupResumeAndAssert()
/// return
/// 									doLookupWithToken()
///																 	doLookupWithToken()
///										rc.LookupRangeDescriptor()
///																	rc.LookupRangeDescriptor()
///										rdc.rangeCacheMu.RLock()
///										rdc.String()
///																	rdc.rangeCacheMu.RLock()
///																	fmt.Printf()
///																	rdc.rangeCacheMu.RUnlock()
///																	rdc.rangeCacheMu.Lock()
///										rdc.rangeCacheMu.RLock()
/// -------------------------------------G2,G3,... deadlock--------------------------------------
```


### backtrace

```
goroutine 34 [semacquire]:
sync.runtime_Semacquire(0xc000014068)
    /usr/local/go/src/runtime/sema.go:56 +0x42
sync.(*WaitGroup).Wait(0xc000014060)
    /usr/local/go/src/sync/waitgroup.go:130 +0x64
command-line-arguments.testRangeCacheCoalescedRquests.func1()
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:55 +0xa8
command-line-arguments.testRangeCacheCoalescedRquests()
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:57 +0x86
created by command-line-arguments.TestCockroach6181
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:62 +0x88

 Goroutine 5 in state semacquire, with sync.runtime_SemacquireMutex on top of the stack:
goroutine 5 [semacquire]:
sync.runtime_SemacquireMutex(0xc00001806c, 0x40a000, 0x0)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*RWMutex).RLock(...)
    /usr/local/go/src/sync/rwmutex.go:50
command-line-arguments.(*rangeDescriptorCache).String(0xc000018060, 0x0, 0x0)
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:31 +0xb0
fmt.(*pp).handleMethods(0xc00010a000, 0x7f8d00000073, 0xc000036601)
    /usr/local/go/src/fmt/print.go:630 +0x302
fmt.(*pp).printArg(0xc00010a000, 0x522e40, 0xc000018060, 0x73)
    /usr/local/go/src/fmt/print.go:713 +0x1e4
fmt.(*pp).doPrintf(0xc00010a000, 0x54c53e, 0x1b, 0xc000110f90, 0x1, 0x1)
    /usr/local/go/src/fmt/print.go:1030 +0x15a
fmt.Fprintf(0x571d40, 0xc0000ac000, 0x54c53e, 0x1b, 0xc000036790, 0x1, 0x1, 0x0, 0x0, 0x0)
    /usr/local/go/src/fmt/print.go:204 +0x72
fmt.Printf(...)
    /usr/local/go/src/fmt/print.go:213
command-line-arguments.(*rangeDescriptorCache).LookupRangeDescriptor(0xc000018060)
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:24 +0xa0
command-line-arguments.doLookupWithToken(...)
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:41
command-line-arguments.testRangeCacheCoalescedRquests.func1.1(0xc00000e010, 0xc000014060)
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:51 +0x2e
created by command-line-arguments.testRangeCacheCoalescedRquests.func1
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:50 +0x8b

 Goroutine 6 in state semacquire, with sync.runtime_SemacquireMutex on top of the stack:
goroutine 6 [semacquire]:
sync.runtime_SemacquireMutex(0xc00001806c, 0x0, 0x0)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*RWMutex).RLock(...)
    /usr/local/go/src/sync/rwmutex.go:50
command-line-arguments.(*rangeDescriptorCache).LookupRangeDescriptor(0xc000018060)
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:23 +0x106
command-line-arguments.doLookupWithToken(...)
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:41
command-line-arguments.testRangeCacheCoalescedRquests.func1.1(0xc00000e010, 0xc000014060)
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:51 +0x2e
created by command-line-arguments.testRangeCacheCoalescedRquests.func1
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:50 +0x8b

 Goroutine 7 in state semacquire, with sync.runtime_SemacquireMutex on top of the stack:
goroutine 7 [semacquire]:
sync.runtime_SemacquireMutex(0xc000018068, 0xc000064000, 0x0)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*RWMutex).Lock(0xc000018060)
    /usr/local/go/src/sync/rwmutex.go:103 +0x88
command-line-arguments.(*rangeDescriptorCache).LookupRangeDescriptor(0xc000018060)
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:26 +0xbf
command-line-arguments.doLookupWithToken(...)
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:41
command-line-arguments.testRangeCacheCoalescedRquests.func1.1(0xc00000e010, 0xc000014060)
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:51 +0x2e
created by command-line-arguments.testRangeCacheCoalescedRquests.func1
    /root/gobench/gobench/goker/blocking/cockroach/6181/cockroach6181_test.go:50 +0x8b
```