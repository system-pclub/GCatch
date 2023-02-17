# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#9935]|[pull request]|[patch]| Blocking | Resource Deadlock | Double Locking |

[cockroach#9935]:(cockroach9935_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/9935/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/9935
 
## Description


This is some description from previous researchers

> This bug is caused by acquiring l.mu.Lock() twice. The fix is
> to release l.mu.Lock() before acquiring l.mu.Lock for the second time.

### backtrace

```
goroutine 21 [semacquire]:
sync.runtime_SemacquireMutex(0xc0000b407c, 0x4d65822107fcfd00, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc0000b4078)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
command-line-arguments.(*loggingT).exit(0xc0000b4078, 0x5733a0, 0xc00004a010)
    /root/gobench/gobench/goker/blocking/cockroach/9935/cockroach9935_test.go:29 +0x78
command-line-arguments.(*loggingT).outputLogEntry(0xc0000b4078)
    /root/gobench/gobench/goker/blocking/cockroach/9935/cockroach9935_test.go:18 +0x85
created by command-line-arguments.TestCockroach9935
    /root/gobench/gobench/goker/blocking/cockroach/9935/cockroach9935_test.go:35 +0xa2
```