
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#82239]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[kubernetes#82239]:(kubernetes82239_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/82239/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/82239
 

## Backtrace

```
Write at 0x00c00001c300 by goroutine 7:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  command-line-arguments.TestKubernetes82239.func1()
      /root/gobench/goker/nonblocking/kubernetes/82239/kubernetes82239_test.go:132 +0x90
  command-line-arguments.TestKubernetes82239()
      /root/gobench/goker/nonblocking/kubernetes/82239/kubernetes82239_test.go:148 +0x275
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Previous read at 0x00c00001c300 by goroutine 9:
  [failed to restore the stack]

Goroutine 7 (running) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1095 +0x537
  testing.runTests.func1()
      /usr/local/go/src/testing/testing.go:1339 +0xa6
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
  testing.runTests()
      /usr/local/go/src/testing/testing.go:1337 +0x594
  testing.(*M).Run()
      /usr/local/go/src/testing/testing.go:1252 +0x2ff
  main.main()
      _testmain.go:44 +0x223

Goroutine 9 (running) created at:
  command-line-arguments.(*PersistentVolumeController).Run()
      /root/gobench/goker/nonblocking/kubernetes/82239/kubernetes82239_test.go:100 +0xb0
```

