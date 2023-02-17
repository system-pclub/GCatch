
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#35501]|[pull request]|[patch]| NonBlocking | Go-Specific | Anonymous function |

[cockroach#35501]:(cockroach35501_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/35501/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/35501
 

## Backtrace

```
Read at 0x00c00004c4a0 by goroutine 12:
  command-line-arguments.validateCheckInTxn()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:19 +0x39
  command-line-arguments.(*SchemaChanger).validateChecks.func1.1()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:59 +0x2b

Previous write at 0x00c00004c4a0 by goroutine 8:
  command-line-arguments.(*SchemaChanger).validateChecks.func1()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:57 +0xa4
  command-line-arguments.(*SchemaChanger).validateChecks()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:62 +0x4c
  command-line-arguments.(*SchemaChanger).runBackfill()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:70 +0x12e
  command-line-arguments.TestCockroach35501.func1()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:79 +0x6c

Goroutine 12 (running) created at:
  command-line-arguments.(*SchemaChanger).validateChecks.func1()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:58 +0xec
  command-line-arguments.(*SchemaChanger).validateChecks()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:62 +0x4c
  command-line-arguments.(*SchemaChanger).runBackfill()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:70 +0x12e
  command-line-arguments.TestCockroach35501.func1()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:79 +0x6c

Goroutine 8 (finished) created at:
  command-line-arguments.TestCockroach35501()
      /root/gobench/goker/nonblocking/cockroach/35501/cockroach35501_test.go:76 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

