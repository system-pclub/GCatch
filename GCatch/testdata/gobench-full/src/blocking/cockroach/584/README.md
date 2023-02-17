# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#584]|[pull request]|[patch]| Blocking | Resource Deadlock | Double Locking |

[cockroach#584]:(cockroach584_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/584/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/584
 
## Description


This is some description from developers

> I'm guessing some of the goroutines might get into deadlock during shutdown. 
> (We cannot use defer as Lock is called inside a for loop.)

Missing unlock before break the loop

### backtrace

```
goroutine 7 [semacquire]:
sync.runtime_SemacquireMutex(0xc0000140f4, 0x0, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc0000140f0)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
command-line-arguments.(*Gossip).manage(0xc0000140f0)
    /root/gobench/gobench/goker/blocking/cockroach/584/cockroach584_test.go:27 +0x6c
command-line-arguments.TestCockroach584.func1(0xc0000140f0)
    /root/gobench/gobench/goker/blocking/cockroach/584/cockroach584_test.go:40 +0x39
created by command-line-arguments.TestCockroach584
    /root/gobench/gobench/goker/blocking/cockroach/584/cockroach584_test.go:36 +0xa6
```