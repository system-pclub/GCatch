
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#4407]|[pull request]|[patch]| NonBlocking | Go-Specific | WaitGroup |

[cockroach#4407]:(cockroach4407_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/4407/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/4407
 

## Backtrace

```
Write at 0x00c00021e570 by goroutine 17:
  internal/race.Write()
      /usr/local/go/src/internal/race/race.go:41 +0x114
  sync.(*WaitGroup).Wait()
      /usr/local/go/src/sync/waitgroup.go:128 +0x115
  command-line-arguments.(*Stopper).Stop()
      /root/gobench/goker/nonblocking/cockroach/4407/cockroach4407_test.go:31 +0x61
  command-line-arguments.TestCockroach4407()
      /root/gobench/goker/nonblocking/cockroach/4407/cockroach4407_test.go:70 +0x1c3
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Previous read at 0x00c00021e570 by goroutine 71:
  internal/race.Read()
      /usr/local/go/src/internal/race/race.go:37 +0x1e8
  sync.(*WaitGroup).Add()
      /usr/local/go/src/sync/waitgroup.go:71 +0x1fb
  command-line-arguments.(*Stopper).RunWorker()
      /root/gobench/goker/nonblocking/cockroach/4407/cockroach4407_test.go:16 +0x47
  command-line-arguments.(*server).Gossip()
      /root/gobench/goker/nonblocking/cockroach/4407/cockroach4407_test.go:44 +0x102

Goroutine 17 (running) created at:
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

Goroutine 71 (finished) created at:
  command-line-arguments.TestCockroach4407()
      /root/gobench/goker/nonblocking/cockroach/4407/cockroach4407_test.go:67 +0x1a2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

