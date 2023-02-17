
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[serving#5865]|[pull request]|[patch]| NonBlocking | Go-Specific | Misuse channel |

[serving#5865]:(serving5865_test.go)
[patch]:https://github.com/ knative/serving/pull/5865/files
[pull request]:https://github.com/ knative/serving/pull/5865
 

## Backtrace

```
panic: send on closed channel [recovered]
	panic: send on closed channel

goroutine 6 [running]:
testing.tRunner.func1.1(0x5d4600, 0x6307e0)
	/usr/local/go/src/testing/testing.go:999 +0x461
testing.tRunner.func1(0xc000122120)
	/usr/local/go/src/testing/testing.go:1002 +0x606
panic(0x5d4600, 0x6307e0)
	/usr/local/go/src/runtime/panic.go:975 +0x3e3
command-line-arguments.(*revisionBackendsManager).endpointsUpdated(...)
	/root/gobench/goker/nonblocking/serving/5865/serving5865_test.go:26
command-line-arguments.TestServing5865.func1(...)
	/root/gobench/goker/nonblocking/serving/5865/serving5865_test.go:50
command-line-arguments.TestServing5865(0xc000122120)
	/root/gobench/goker/nonblocking/serving/5865/serving5865_test.go:51 +0xaa
testing.tRunner(0xc000122120, 0x60af80)
	/usr/local/go/src/testing/testing.go:1050 +0x1ec
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1095 +0x538
```

