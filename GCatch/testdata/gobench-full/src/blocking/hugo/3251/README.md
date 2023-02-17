
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[hugo#3251]|[pull request]|[patch]| Blocking | Resource Deadlock | AB-BA deadlock |

[hugo#3251]:(hugo3251_test.go)
[patch]:https://github.com/gohugoio/hugo/pull/3251/files
[pull request]:https://github.com/gohugoio/hugo/pull/3251
 
## Description

A goroutine can hold Lock() at line 20 then acquire RLock() at 
line 29. RLock() at line 29 will never be acquired because Lock() 
at line 20 will never be released. 

## Backtrace

```
goroutine 9 [semacquire]:
sync.runtime_SemacquireMutex(0xc000096004, 0xc000182000, 0x1)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc000096000)
	/usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
	/usr/local/go/src/sync/mutex.go:81
command-line-arguments.(*remoteLock).URLLock(0x645cc0, 0x54c311, 0x1a)
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:25 +0xda
command-line-arguments.resGetRemote(0x54c311, 0x1a, 0x0, 0x0)
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:38 +0x5f
command-line-arguments.TestHugo3251.func1(0xc000014100, 0x54c311, 0x1a, 0x2)
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:54 +0x8f
created by command-line-arguments.TestHugo3251
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:51 +0xdd

goroutine 16 [semacquire]:
sync.runtime_SemacquireMutex(0x645cc4, 0x539600, 0x1)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0x645cc0)
	/usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
	/usr/local/go/src/sync/mutex.go:81
sync.(*RWMutex).Lock(0x645cc0)
	/usr/local/go/src/sync/rwmutex.go:98 +0x97
command-line-arguments.(*remoteLock).URLLock(0x645cc0, 0x54c311, 0x1a)
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:21 +0x31
command-line-arguments.resGetRemote(0x54c311, 0x1a, 0x0, 0x0)
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:38 +0x5f
command-line-arguments.TestHugo3251.func1(0xc000014100, 0x54c311, 0x1a, 0x9)
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:54 +0x8f
created by command-line-arguments.TestHugo3251
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:51 +0xdd

goroutine 17 [semacquire]:
sync.runtime_SemacquireMutex(0x645ccc, 0xc000098000, 0x0)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*RWMutex).RLock(...)
	/usr/local/go/src/sync/rwmutex.go:50
command-line-arguments.(*remoteLock).URLUnlock(0x645cc0, 0x54c311, 0x1a)
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:30 +0xf2
command-line-arguments.resGetRemote.func1(0x54c311, 0x1a)
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:39 +0x41
command-line-arguments.resGetRemote(0x54c311, 0x1a, 0x0, 0x0)
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:41 +0xa9
command-line-arguments.TestHugo3251.func1(0xc000014100, 0x54c311, 0x1a, 0xa)
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:54 +0x8f
created by command-line-arguments.TestHugo3251
	/root/gobench/goker/blocking/hugo/3251/hugo3251_test.go:51 +0xdd
```

