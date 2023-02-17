
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[serving#6171]|[pull request]|[patch]| NonBlocking | Go-Specific | Testing library |

[serving#6171]:(serving6171_test.go)
[patch]:https://github.com/knative/serving/pull/6171/files
[pull request]:https://github.com/knative/serving/pull/6171
 

## Backtrace

```
panic: Log in goroutine after TestServing6171/Serving6171 has completed

goroutine 11 [running]:
testing.(*common).logDepth(0xc000122240, 0x73a271, 0x1, 0x3)
	/usr/local/go/src/testing/testing.go:738 +0x7c4
testing.(*common).log(...)
	/usr/local/go/src/testing/testing.go:720
testing.(*common).Logf(0xc000122240, 0x602ce8, 0x2, 0xc00004c540, 0x1, 0x1)
	/usr/local/go/src/testing/testing.go:766 +0x90
command-line-arguments.testingWriter.Write(...)
	/root/gobench/goker/nonblocking/serving/6171/serving6171_test.go:36
command-line-arguments.(*ioCore).Write(0xc00004c4d0)
	/root/gobench/goker/nonblocking/serving/6171/serving6171_test.go:74 +0x51
command-line-arguments.(*CheckedEntry).Write(0xc000049f88)
	/root/gobench/goker/nonblocking/serving/6171/serving6171_test.go:23 +0x92
command-line-arguments.(*SugaredLogger).log(0xc000010038)
	/root/gobench/goker/nonblocking/serving/6171/serving6171_test.go:97 +0x17d
command-line-arguments.(*SugaredLogger).Errorw(...)
	/root/gobench/goker/nonblocking/serving/6171/serving6171_test.go:101
command-line-arguments.(*revisionWatcher).checkDests.func1(0xc000010040)
	/root/gobench/goker/nonblocking/serving/6171/serving6171_test.go:120 +0x4c
created by command-line-arguments.(*revisionWatcher).checkDests
	/root/gobench/goker/nonblocking/serving/6171/serving6171_test.go:119 +0x4d
```

