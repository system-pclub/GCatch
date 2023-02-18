
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#7443]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[etcd#7443]:(etcd7443_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/7443/files
[pull request]:https://github.com/etcd-io/etcd/pull/7443
 
## Description

Line 161 holds mutex at 158 but try to send messages to `notifyCh` at 161.
`Close()` goroutine blocked at line 172 for acquiring the same mutex. 


## Backtrace

```
goroutine 566 [chan receive, 10 minutes]:
command-line-arguments.TestEtcd7443(0xc0002a5320)
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:214 +0x184
testing.tRunner(0xc0002a5320, 0x5537b8)
	/usr/local/go/src/testing/testing.go:1050 +0xdc
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1095 +0x28b

goroutine 567 [chan send, 10 minutes]:
command-line-arguments.(*simpleBalancer).Up.func2()
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:162 +0x85
command-line-arguments.(*addrConn).tearDown(0xc000353f50)
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:23 +0x86
command-line-arguments.(*ClientConn).lbWatcher(0xc000353ef0)
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:73 +0x3ac
created by command-line-arguments.DialContext
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:125 +0xab

goroutine 570 [semacquire, 10 minutes]:
sync.runtime_SemacquireMutex(0xc000074d94, 0x1, 0x1)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc000074d90)
	/usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
	/usr/local/go/src/sync/mutex.go:81
sync.(*RWMutex).Lock(0xc000074d90)
	/usr/local/go/src/sync/rwmutex.go:98 +0x97
command-line-arguments.(*simpleBalancer).Up(0xc000074d70, 0x2, 0x0)
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:146 +0x50
command-line-arguments.(*addrConn).resetTransport(0xc000353f80)
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:31 +0x84
command-line-arguments.(*ClientConn).resetAddrConn.func1(0xc000353f80)
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:92 +0x2b
created by command-line-arguments.(*ClientConn).resetAddrConn
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:91 +0x105

goroutine 571 [semacquire, 10 minutes]:
sync.runtime_SemacquireMutex(0xc000074d94, 0x0, 0x1)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc000074d90)
	/usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
	/usr/local/go/src/sync/mutex.go:81
sync.(*RWMutex).Lock(0xc000074d90)
	/usr/local/go/src/sync/rwmutex.go:98 +0x97
command-line-arguments.(*simpleBalancer).Close(0xc000074d70)
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:173 +0x47
command-line-arguments.TestEtcd7443.func1(0xc000375320, 0xc000074d70)
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:211 +0x53
created by command-line-arguments.TestEtcd7443
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:209 +0x14b

goroutine 572 [semacquire, 10 minutes]:
sync.runtime_SemacquireMutex(0xc000074d94, 0x0, 0x1)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc000074d90)
	/usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
	/usr/local/go/src/sync/mutex.go:81
sync.(*RWMutex).Lock(0xc000074d90)
	/usr/local/go/src/sync/rwmutex.go:98 +0x97
command-line-arguments.(*simpleBalancer).Close(0xc000074d70)
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:173 +0x47
command-line-arguments.(*ClientConn).Close(0xc000353ef0)
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:102 +0x119
created by command-line-arguments.TestEtcd7443
	/root/gobench/goker/blocking/etcd/7443/etcd7443_test.go:213 +0x16d
```

