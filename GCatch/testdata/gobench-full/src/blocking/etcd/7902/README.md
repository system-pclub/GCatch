
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#7902]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[etcd#7902]:(etcd7902_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/7902/files
[pull request]:https://github.com/etcd-io/etcd/pull/7902
 
## Description

Some description from developers or pervious reseachers

>  At least two goroutines are needed to trigger this bug,
>  one is leader and the other is follower. Both the leader 
>  and the follower execute the code above. If the follower
>  acquires mu.Lock() firstly and enter rc.release(), it will
>  be blocked at <- rcNextc (nextc). Only the leader can execute 
>  close(nextc) to unblock the follower inside rc.release().
>  However, in order to invoke rc.release(), the leader needs
>  to acquires mu.Lock(). 
>  The fix is to remove the lock and unlock around rc.release().

Possible intervening

```
///
/// G1						G2 (leader)					G3 (follower)
/// runElectionFunc()
/// doRounds()
/// wg.Wait()
/// 						...
/// 						mu.Lock()
/// 						rc.validate()
/// 						rcNextc = nextc
/// 						mu.Unlock()					...
/// 													mu.Lock()
/// 													rc.validate()
/// 													mu.Unlock()
/// 													mu.Lock()
/// 													rc.release()
/// 													<-rcNextc
/// 						mu.Lock()
/// -------------------------G1,G2,G3 deadlock--------------------------
///
```

## Backtrace

```
goroutine 19 [semacquire]:
sync.runtime_Semacquire(0xc000014088)
    /usr/local/go/src/runtime/sema.go:56 +0x42
sync.(*WaitGroup).Wait(0xc000014080)
    /usr/local/go/src/sync/waitgroup.go:130 +0x64
command-line-arguments.doRounds(0xc00006a060, 0x3, 0x3, 0x64)
    /root/gobench/goker/blocking/etcd/7902/etcd7902_test.go:59 +0xd8
command-line-arguments.runElectionFunc()
    /root/gobench/goker/blocking/etcd/7902/etcd7902_test.go:37 +0x326
created by command-line-arguments.TestEtcd7902
    /root/gobench/goker/blocking/etcd/7902/etcd7902_test.go:64 +0x88

 Goroutine 5 in state semacquire, with sync.runtime_SemacquireMutex on top of the stack:
goroutine 5 [semacquire]:
sync.runtime_SemacquireMutex(0xc00001407c, 0x551c00, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc000014078)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
command-line-arguments.doRounds.func1(0xc000014080, 0x64, 0xc000014078, 0xc00006a060)
    /root/gobench/goker/blocking/etcd/7902/etcd7902_test.go:53 +0x126
created by command-line-arguments.doRounds
    /root/gobench/goker/blocking/etcd/7902/etcd7902_test.go:44 +0xb8

 Goroutine 6 in state chan receive, with command-line-arguments.runElectionFunc.func4 on top of the stack:
goroutine 6 [chan receive]:
command-line-arguments.runElectionFunc.func4()
    /root/gobench/goker/blocking/etcd/7902/etcd7902_test.go:34 +0x48
command-line-arguments.doRounds.func1(0xc000014080, 0x64, 0xc000014078, 0xc00006a080)
    /root/gobench/goker/blocking/etcd/7902/etcd7902_test.go:54 +0xed
created by command-line-arguments.doRounds
    /root/gobench/goker/blocking/etcd/7902/etcd7902_test.go:44 +0xb8

 Goroutine 7 in state semacquire, with sync.runtime_SemacquireMutex on top of the stack:
goroutine 7 [semacquire]:
sync.runtime_SemacquireMutex(0xc00001407c, 0x551c00, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc000014078)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
command-line-arguments.doRounds.func1(0xc000014080, 0x64, 0xc000014078, 0xc00006a0a0)
    /root/gobench/goker/blocking/etcd/7902/etcd7902_test.go:53 +0x126
created by command-line-arguments.doRounds
    /root/gobench/goker/blocking/etcd/7902/etcd7902_test.go:44 +0xb8
```

