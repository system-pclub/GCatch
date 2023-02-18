
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#27037]|[pull request]|[patch]| NonBlocking | Go-Specific | Anonymous function |

[moby#27037]:(moby27037_test.go)
[patch]:https://github.com/moby/moby/pull/27037/files
[pull request]:https://github.com/moby/moby/pull/27037
 

## Backtrace

```
Read at 0x00c0000a6088 by goroutine 9:
  command-line-arguments.TestMoby27037.func1()
      /root/gobench/goker/nonblocking/moby/27037/moby27037_test.go:15 +0x8a

Previous write at 0x00c0000a6088 by goroutine 7:
  command-line-arguments.TestMoby27037()
      /root/gobench/goker/nonblocking/moby/27037/moby27037_test.go:11 +0x10a
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 9 (running) created at:
  command-line-arguments.TestMoby27037()
      /root/gobench/goker/nonblocking/moby/27037/moby27037_test.go:13 +0xe6
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

```
Read at 0x00c0000a6088 by goroutine 8:
  command-line-arguments.TestMoby27037.func1()
      /root/gobench/goker/nonblocking/moby/27037/moby27037_test.go:15 +0x8a

Previous write at 0x00c0000a6088 by goroutine 7:
  command-line-arguments.TestMoby27037()
      /root/gobench/goker/nonblocking/moby/27037/moby27037_test.go:11 +0x10a
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Goroutine 8 (running) created at:
  command-line-arguments.TestMoby27037()
      /root/gobench/goker/nonblocking/moby/27037/moby27037_test.go:13 +0xe6
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
