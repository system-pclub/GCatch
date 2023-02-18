# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#3710]|[pull request]|[patch]| Blocking | Resource Deadlock | RWR Deadlock |

[cockroach#3710]:(cockroach3710_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/3710/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/3710
 
## Description


This is some description from previous researchers

> This deadlock is casued by acquiring a RLock twice in a call chain.
> ForceRaftLogScanAndProcess(acquire s.mu.RLock()) ->MaybeAdd()->shouldQueue()->
> getTruncatableIndexes()->RaftStatus(acquire s.mu.Rlock())

Possible intervening

```
/// G1 										G2
/// store.ForceRaftLogScanAndProcess()
/// s.mu.RLock()
/// s.raftLogQueue.MaybeAdd()
/// bq.impl.shouldQueue()
/// getTruncatableIndexes()
/// r.store.RaftStatus()
/// 										store.processRaft()
/// 										s.mu.Lock()
/// s.mu.RLock()
/// ----------------------G1,G2 deadlock---------------------
```

### backtrace

```
goroutine 12205 [semacquire]:
sync.runtime_SemacquireMutex(0xc00009bbdc, 0x0, 0x0)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*RWMutex).RLock(...)
    /usr/local/go/src/sync/rwmutex.go:50
command-line-arguments.(*Store).RaftStatus(0xc00009bbc0)
    /root/gobench/gobench/goker/blocking/cockroach/3710/cockroach3710_test.go:28 +0x92
command-line-arguments.getTruncatableIndexes(...)
    /root/gobench/gobench/goker/blocking/cockroach/3710/cockroach3710_test.go:68
command-line-arguments.(*raftLogQueue).shouldQueue(...)
    /root/gobench/gobench/goker/blocking/cockroach/3710/cockroach3710_test.go:64
command-line-arguments.(*baseQueue).MaybeAdd(0xc0000965f0, 0xc0004c0058)
    /root/gobench/gobench/goker/blocking/cockroach/3710/cockroach3710_test.go:58 +0x6e
command-line-arguments.(*Store).ForceRaftLogScanAndProcess(0xc00009bbc0)
    /root/gobench/gobench/goker/blocking/cockroach/3710/cockroach3710_test.go:22 +0xa2
created by command-line-arguments.TestCockroach3710
    /root/gobench/gobench/goker/blocking/cockroach/3710/cockroach3710_test.go:93 +0x9b

 Goroutine 12036 in state semacquire, with sync.runtime_SemacquireMutex on top of the stack:
goroutine 12036 [semacquire]:
sync.runtime_SemacquireMutex(0xc00009bbd8, 0xc00009b000, 0x0)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*RWMutex).Lock(0xc00009bbd0)
    /usr/local/go/src/sync/rwmutex.go:103 +0x88
command-line-arguments.(*Store).processRaft.func1(0xc00009bbc0)
    /root/gobench/gobench/goker/blocking/cockroach/3710/cockroach3710_test.go:36 +0x4b
created by command-line-arguments.(*Store).processRaft
    /root/gobench/gobench/goker/blocking/cockroach/3710/cockroach3710_test.go:33 +0x3f
```