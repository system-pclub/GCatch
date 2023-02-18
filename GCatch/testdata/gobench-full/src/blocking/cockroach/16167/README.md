# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#16167]|[pull request]|[patch]| Blocking | Resource Deadlock | Double Locking |

[cockroach#16167]:(cockroach16167_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/16167/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/16167
 
## Description


This is some description from previous researchers

> This is another example for deadlock caused by recursively
> acquiring RWLock. There are two lock variables (systemConfigCond and systemConfigMu)
> involved in this bug, but they are actually the same lock, which can be found from
> the following code.
> There are two goroutine involved in this deadlock. The first goroutine acquires
> systemConfigMu.Lock() firstly, then tries to acquire systemConfigMu.RLock(). The
> second goroutine tries to acquire systemConfigMu.Lock(). If the second goroutine
> interleaves in between the two lock operations of the first goroutine, deadlock will happen.

Possible intervening

```
/// G1 							G2
/// e.Start()
/// e.updateSystemConfig()
/// 							e.execParsed()
/// 							e.systemConfigCond.L.Lock()
/// e.systemConfigMu.Lock()
/// 							e.systemConfigMu.RLock()
/// ----------------------G1,G2 deadlock--------------------
```


### backtrace

```
goroutine 44794 [semacquire]:
sync.runtime_SemacquireMutex(0xc0003f7874, 0x7fdde274ae00, 0x0)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*RWMutex).RLock(...)
	/usr/local/go/src/sync/rwmutex.go:50
command-line-arguments.(*Executor).getDatabaseCache(0xc0003f7860)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:69 +0x91
command-line-arguments.(*Session).resetForBatch(...)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:38
command-line-arguments.(*Executor).Prepare(...)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:65
command-line-arguments.PreparedStatements.New(...)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:30
command-line-arguments.(*Executor).execStmtInOpenTxn(...)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:61
command-line-arguments.(*Executor).execStmtsInCurrentTxn(...)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:57
command-line-arguments.runTxnAttempt(0xc0003f7860, 0xc0000aae28)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:79 +0x35
command-line-arguments.(*Executor).execParsed(0xc0003f7860, 0xc0000aae28)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:53 +0x7e
command-line-arguments.TestCockroach16167(0xc000482d80)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:101 +0xd2
testing.tRunner(0xc000482d80, 0x54d0d8)
	/usr/local/go/src/testing/testing.go:1123 +0xef
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1168 +0x2b3

goroutine 44795 [semacquire]:
sync.runtime_SemacquireMutex(0xc0003f7870, 0x0, 0x0)
	/usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*RWMutex).Lock(0xc0003f7868)
	/usr/local/go/src/sync/rwmutex.go:103 +0x85
command-line-arguments.(*Executor).updateSystemConfig(0xc0003f7860)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:74 +0x45
command-line-arguments.(*Executor).Start(0xc0003f7860)
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:47 +0x2b
created by command-line-arguments.TestCockroach16167
	/home/yuanting/work-gobench/gobench/gobench/goker/blocking/cockroach/16167/cockroach16167_test.go:100 +0xba
```