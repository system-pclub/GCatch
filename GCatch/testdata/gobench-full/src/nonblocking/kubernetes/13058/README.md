
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#13058]|[pull request]|[patch]| NonBlocking | Go-Specific | WaitGroup |

[kubernetes#13058]:(kubernetes13058_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/13058/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/13058
 

## Backtrace

```
panic: sync: negative WaitGroup counter

goroutine 30781 [running]:
sync.(*WaitGroup).Add(0xc0004a7a60, 0xffffffffffffffff)
	/usr/local/go/src/sync/waitgroup.go:74 +0x2c0
sync.(*WaitGroup).Done(...)
	/usr/local/go/src/sync/waitgroup.go:99
command-line-arguments.TestKubernetes13058.func1(0x0, 0x0)
	/root/gobench/goker/nonblocking/kubernetes/13058/kubernetes13058_test.go:78 +0x4c
command-line-arguments.ResourceEventHandlerFuncs.OnDelete(0xc000497090, 0x0, 0x0)
	/root/gobench/goker/nonblocking/kubernetes/13058/kubernetes13058_test.go:26 +0x5c
command-line-arguments.NewInformer.func1(0x0, 0x0)
	/root/gobench/goker/nonblocking/kubernetes/13058/kubernetes13058_test.go:53 +0x5b
command-line-arguments.(*Controller).processLoop(0xc0004da190)
	/root/gobench/goker/nonblocking/kubernetes/13058/kubernetes13058_test.go:36 +0x4f
command-line-arguments.Until.func1(...)
	/root/gobench/goker/nonblocking/kubernetes/13058/kubernetes13058_test.go:67
command-line-arguments.Until(0xc00042efb0, 0x989680, 0xc0004dd500)
	/root/gobench/goker/nonblocking/kubernetes/13058/kubernetes13058_test.go:68 +0x38
command-line-arguments.(*Controller).Run(0xc0004da190, 0xc0004dd500)
	/root/gobench/goker/nonblocking/kubernetes/13058/kubernetes13058_test.go:42 +0x6a
created by command-line-arguments.TestKubernetes13058
	/root/gobench/goker/nonblocking/kubernetes/13058/kubernetes13058_test.go:83 +0x151
```

