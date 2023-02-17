
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[serving#3068]|[pull request]|[patch]| NonBlocking | Go-Specific | Misuse channel |

[serving#3068]:(serving3068_test.go)
[patch]:https://github.com/ knative/serving/pull/3068/files
[pull request]:https://github.com/ knative/serving/pull/3068
 

## Backtrace

```
panic: send on closed channel

goroutine 21 [running]:
command-line-arguments.(*impl).Go(0xc000116300, 0xc000112480)
	/root/gobench/goker/nonblocking/serving/3068/serving3068_test.go:44 +0x7c
command-line-arguments.TestServing3068.func1(0x6341a0, 0xc000116300, 0xc00011e088, 0xc00011e090)
	/root/gobench/goker/nonblocking/serving/3068/serving3068_test.go:65 +0x4c
created by command-line-arguments.TestServing3068
	/root/gobench/goker/nonblocking/serving/3068/serving3068_test.go:63 +0x11e
```
