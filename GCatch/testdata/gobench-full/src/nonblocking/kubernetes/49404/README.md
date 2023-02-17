
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#49404]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[kubernetes#49404]:(kubernetes49404_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/49404/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/49404
 

## Backtrace

```
Write at 0x00c000126088 by goroutine 9:
  command-line-arguments.TestKubernetes49404.func1.1()
      /root/gobench/goker/nonblocking/kubernetes/49404/kubernetes49404_test.go:130 +0x46
  command-line-arguments.websocket_Handler.ServeHTTP()
      /root/gobench/goker/nonblocking/kubernetes/49404/kubernetes49404_test.go:17 +0x34
  command-line-arguments.(*ServeMux).ServeHTTP()
      /root/gobench/goker/nonblocking/kubernetes/49404/kubernetes49404_test.go:51 +0x4e
  command-line-arguments.serverHandler.ServeHTTP()
      /root/gobench/goker/nonblocking/kubernetes/49404/kubernetes49404_test.go:101 +0x68
  command-line-arguments.(*conn).serve()
      /root/gobench/goker/nonblocking/kubernetes/49404/kubernetes49404_test.go:92 +0x2b

Previous read at 0x00c000126088 by goroutine 7:
  command-line-arguments.TestKubernetes49404.func1.2()
      /root/gobench/goker/nonblocking/kubernetes/49404/kubernetes49404_test.go:138 +0x38
  command-line-arguments.TestKubernetes49404.func1()
      /root/gobench/goker/nonblocking/kubernetes/49404/kubernetes49404_test.go:142 +0x2f1
  command-line-arguments.TestKubernetes49404()
      /root/gobench/goker/nonblocking/kubernetes/49404/kubernetes49404_test.go:142 +0x63
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 9 (running) created at:
  command-line-arguments.(*http_Server).Serve()
      /root/gobench/goker/nonblocking/kubernetes/49404/kubernetes49404_test.go:110 +0x85
  command-line-arguments.(*Server).goServe.func1()
      /root/gobench/goker/nonblocking/kubernetes/49404/kubernetes49404_test.go:83 +0x83

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
```

