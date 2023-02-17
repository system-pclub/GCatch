
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#7492]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[etcd#7492]:(etcd7492_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/7492/files
[pull request]:https://github.com/etcd-io/etcd/pull/7492
 
## Description

Possible intervening

```
///
///	G1										G2
///											stk.run()
///	ts.assignSimpleTokenToUser()
///	t.simpleTokensMu.Lock()
///	t.simpleTokenKeeper.addSimpleToken()
///	tm.addSimpleTokenCh <- true
///											<-tm.addSimpleTokenCh
///	t.simpleTokensMu.Unlock()
///	ts.assignSimpleTokenToUser()
///	...										...
///	t.simpleTokensMu.Lock()
///											<-tokenTicker.C
///	tm.addSimpleTokenCh <- true
///											tm.deleteTokenFunc()
///											t.simpleTokensMu.Lock()
///------------------------------------G1,G2 deadlock---------------------------------------------
///
```

See the [real bug](../../../../goreal/blocking/etcd/7492/README.md)

## Backtrace

```
goroutine 1077 [semacquire, 9 minutes]:
sync.runtime_Semacquire(0xc0002461a8)
	/usr/local/go/src/runtime/sema.go:56 +0x42
sync.(*WaitGroup).Wait(0xc0002461a0)
	/usr/local/go/src/sync/waitgroup.go:130 +0x64
command-line-arguments.TestEtcd7492(0xc00024ca20)
	/root/gobench/goker/blocking/etcd/7492/etcd7492_test.go:134 +0x120
testing.tRunner(0xc00024ca20, 0x554708)
	/usr/local/go/src/testing/testing.go:1050 +0xdc
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1095 +0x28b

goroutine 1080 [chan send, 9 minutes]:
command-line-arguments.(*simpleTokenTTLKeeper).addSimpleToken(...)
	/root/gobench/goker/blocking/etcd/7492/etcd7492_test.go:63
command-line-arguments.(*tokenSimple).assignSimpleTokenToUser(0xc000098240)
	/root/gobench/goker/blocking/etcd/7492/etcd7492_test.go:84 +0x55
command-line-arguments.(*tokenSimple).assign(0xc000098240)
	/root/gobench/goker/blocking/etcd/7492/etcd7492_test.go:79 +0x2b
command-line-arguments.(*authStore).Authenticate(...)
	/root/gobench/goker/blocking/etcd/7492/etcd7492_test.go:28
command-line-arguments.TestEtcd7492.func1(0xc0002461a0, 0xc00005e650)
	/root/gobench/goker/blocking/etcd/7492/etcd7492_test.go:131 +0x5c
created by command-line-arguments.TestEtcd7492
	/root/gobench/goker/blocking/etcd/7492/etcd7492_test.go:129 +0x104

goroutine 1078 [semacquire, 9 minutes]:
sync.runtime_SemacquireMutex(0xc00009824c, 0xc000073e01, 0x1)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc000098248)
	/usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
	/usr/local/go/src/sync/mutex.go:81
sync.(*RWMutex).Lock(0xc000098248)
	/usr/local/go/src/sync/rwmutex.go:98 +0x97
command-line-arguments.newDeleterFunc.func1(0x54a981, 0x1)
	/root/gobench/goker/blocking/etcd/7492/etcd7492_test.go:89 +0x42
command-line-arguments.(*simpleTokenTTLKeeper).run(0xc000098260)
	/root/gobench/goker/blocking/etcd/7492/etcd7492_test.go:52 +0xa2
created by command-line-arguments.NewSimpleTokenTTLKeeper
	/root/gobench/goker/blocking/etcd/7492/etcd7492_test.go:38 +0xd8
```

