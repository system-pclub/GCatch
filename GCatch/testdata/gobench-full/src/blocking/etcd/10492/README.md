
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#10492]|[pull request]|[patch]| Blocking | Resource Deadlock | Double locking |

[etcd#10492]:(etcd10492_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/10492/files
[pull request]:https://github.com/etcd-io/etcd/pull/10492
 
## Description

line 19, 31 double locking

## Backtrace

```
goroutine 6 [semacquire]:
sync.runtime_SemacquireMutex(0xc0000782d4, 0x0, 0x1)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc0000782d0)
	/usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
	/usr/local/go/src/sync/mutex.go:81
sync.(*RWMutex).Lock(0xc0000782d0)
	/usr/local/go/src/sync/rwmutex.go:98 +0x97
command-line-arguments.(*lessor).Checkpoint(0xc0000782d0)
	/root/gobench/goker/blocking/etcd/10492/etcd10492_test.go:20 +0x3a
command-line-arguments.TestEtcd10492.func1(0x572160, 0xc000014080)
	/root/gobench/goker/blocking/etcd/10492/etcd10492_test.go:46 +0x2a
command-line-arguments.(*lessor).Renew(0xc0000782d0)
	/root/gobench/goker/blocking/etcd/10492/etcd10492_test.go:37 +0xb9
command-line-arguments.TestEtcd10492(0xc000122120)
	/root/gobench/goker/blocking/etcd/10492/etcd10492_test.go:51 +0xf7
testing.tRunner(0xc000122120, 0x550738)
	/usr/local/go/src/testing/testing.go:1050 +0xdc
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1095 +0x28b
```

