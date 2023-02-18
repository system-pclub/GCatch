
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#18412]|[pull request]|[patch]| NonBlocking | Traditional | Order violation |

[moby#18412]:(moby18412_test.go)
[patch]:https://github.com/moby/moby/pull/18412/files
[pull request]:https://github.com/moby/moby/pull/18412
 

## Backtrace

```
Read at 0x00c000100048 by goroutine 7:
  bytes.(*Buffer).String()
      /usr/local/go/src/bytes/buffer.go:65 +0x64d
  command-line-arguments.RunCommandWithOutputForDuration()
      /root/gobench/goker/nonblocking/moby/18412/moby18412_test.go:52 +0x5ab
  command-line-arguments.TestMoby18412()
      /root/gobench/goker/nonblocking/moby/18412/moby18412_test.go:58 +0xb1
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb

Previous write at 0x00c000100048 by goroutine 8:
  bytes.(*Buffer).grow()
      /usr/local/go/src/bytes/buffer.go:147 +0x27f
  bytes.(*Buffer).ReadFrom()
      /usr/local/go/src/bytes/buffer.go:202 +0x7c
  io.copyBuffer()
      /usr/local/go/src/io/io.go:391 +0x3fa
  io.Copy()
      /usr/local/go/src/io/io.go:364 +0x7a
  os/exec.(*Cmd).writerDescriptor.func1()
      /usr/local/go/src/os/exec/exec.go:311 +0x4a
  os/exec.(*Cmd).Start.func1()
      /usr/local/go/src/os/exec/exec.go:441 +0x34

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

Goroutine 8 (running) created at:
  os/exec.(*Cmd).Start()
      /usr/local/go/src/os/exec/exec.go:440 +0xa9d
  command-line-arguments.RunCommandWithOutputForDuration()
      /root/gobench/goker/nonblocking/moby/18412/moby18412_test.go:29 +0x360
  command-line-arguments.TestMoby18412()
      /root/gobench/goker/nonblocking/moby/18412/moby18412_test.go:58 +0xb1
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

