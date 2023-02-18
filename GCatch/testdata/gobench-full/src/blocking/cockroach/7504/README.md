# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#7504]|[pull request]|[patch]| Blocking | Resource Deadlock | AB-BA Deadlock |

[cockroach#7504]:(cockroach7504_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/7504/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/7504
 
## Description


This is some description from previous researchers

> There are locking leaseState, tableNameCache in Release(), but
> tableNameCache,LeaseState in AcquireByName.  It is AB and BA deadlock.


### backtrace

```
goroutine 49274 [semacquire]:
sync.runtime_SemacquireMutex(0xc000102814, 0x0, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc000102810)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
command-line-arguments.(*tableNameCache).remove(0xc000102810, 0xc000014510)
    /root/gobench/gobench/goker/blocking/cockroach/7504/cockroach7504_test.go:82 +0x112
command-line-arguments.(*tableState).removeLease(0xc00000c3a0, 0xc000014510)
    /root/gobench/gobench/goker/blocking/cockroach/7504/cockroach7504_test.go:57 +0x54
command-line-arguments.(*tableState).release(0xc00000c3a0, 0xc000014510)
    /root/gobench/gobench/goker/blocking/cockroach/7504/cockroach7504_test.go:53 +0xb6
command-line-arguments.(*LeaseManager).Release(0xc000060a80, 0xc000014510)
    /root/gobench/gobench/goker/blocking/cockroach/7504/cockroach7504_test.go:116 +0x70
command-line-arguments.TestCockroach7504.func2(0xc000060a80, 0xc00000c380)
    /root/gobench/gobench/goker/blocking/cockroach/7504/cockroach7504_test.go:161 +0x44
created by command-line-arguments.TestCockroach7504
    /root/gobench/gobench/goker/blocking/cockroach/7504/cockroach7504_test.go:159 +0x27a

 Goroutine 49273 in state semacquire, with sync.runtime_SemacquireMutex on top of the stack:
goroutine 49273 [semacquire]:
sync.runtime_SemacquireMutex(0xc000014514, 0xc000102300, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc000014510)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
command-line-arguments.(*tableNameCache).get(0xc000102810, 0x0)
    /root/gobench/gobench/goker/blocking/cockroach/7504/cockroach7504_test.go:76 +0x103
command-line-arguments.(*LeaseManager).AcquireByName(...)
    /root/gobench/gobench/goker/blocking/cockroach/7504/cockroach7504_test.go:103
command-line-arguments.TestCockroach7504.func1(0xc000060a80)
    /root/gobench/gobench/goker/blocking/cockroach/7504/cockroach7504_test.go:156 +0x38
created by command-line-arguments.TestCockroach7504
    /root/gobench/gobench/goker/blocking/cockroach/7504/cockroach7504_test.go:154 +0x24e
```