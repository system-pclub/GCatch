
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#22941]|[pull request]|[patch]| NonBlocking | Go-Specific | Anonymous function |

[moby#22941]:(moby22941_test.go)
[patch]:https://github.com/moby/moby/pull/22941/files
[pull request]:https://github.com/moby/moby/pull/22941
 

## Backtrace

```
Read at 0x00c00009e2d0 by goroutine 8:
  command-line-arguments.TestMoby22941.func1()
      /root/gobench/goker/nonblocking/moby/22941/moby22941_test.go:41 +0x46

Previous write at 0x00c00009e2d0 by goroutine 7:
  command-line-arguments.TestMoby22941()
      /root/gobench/goker/nonblocking/moby/22941/moby22941_test.go:39 +0x33c
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 8 (running) created at:
  command-line-arguments.TestMoby22941()
      /root/gobench/goker/nonblocking/moby/22941/moby22941_test.go:40 +0x39e
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

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

