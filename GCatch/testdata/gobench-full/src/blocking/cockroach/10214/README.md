# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#10214]|[pull request]|[patch]| Blocking | Resource Deadlock | AB-BA deadlock |

[cockroach#10214]:(cockroach10214_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/10214/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/10214
 
## Description


This is some description from previous researchers

> This deadlock is caused by different order when acquiring
> coalescedMu.Lock() and raftMu.Lock(). The fix is to refactor sendQueuedHeartbeats()
> so that cockroachdb can unlock coalescedMu before locking raftMu.

### backtrace

```
goroutine 25176 [semacquire]:
sync.runtime_SemacquireMutex(0xc00009a7b4, 0x0, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc00009a7b0)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
command-line-arguments.(*Replica).maybeCoalesceHeartbeat(0xc00000c320, 0xc00000c4e0)
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:74 +0x99
command-line-arguments.(*Replica).maybeQuiesceLocked(0xc00000c320, 0xc00001b8f0)
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:64 +0x43
command-line-arguments.(*Replica).tickRaftMuLocked(0xc00000c320)
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:58 +0x69
command-line-arguments.(*Replica).tick(0xc00000c320)
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:51 +0x64
command-line-arguments.TestCockroach10214.func2(0xc00000c320)
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:103 +0x2b
created by command-line-arguments.TestCockroach10214
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:102 +0x217

 Goroutine 25175 in state semacquire, with sync.runtime_SemacquireMutex on top of the stack:
goroutine 25175 [semacquire]:
sync.runtime_SemacquireMutex(0xc00000c324, 0x0, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc00000c320)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
command-line-arguments.(*Replica).reportUnreachable(0xc00000c320)
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:42 +0x78
command-line-arguments.(*Store).sendQueuedHeartbeatsToNode(0xc00009a7b0)
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:31 +0x56
command-line-arguments.(*Store).sendQueuedHeartbeats(0xc00009a7b0)
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:24 +0x6d
command-line-arguments.TestCockroach10214.func1(0xc00009a7b0)
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:99 +0x2b
created by command-line-arguments.TestCockroach10214
    /root/gobench/gobench/goker/blocking/cockroach/10214/cockroach10214_test.go:98 +0x1f5
```