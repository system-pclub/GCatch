# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#25456]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[cockroach#25456]:(cockroach25456_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/25456/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/25456
 
## Description


This is some description from developers

> When CheckConsistency returns an error, the queue checks whether the
  store is draining to decide whether the error is worth logging.
  Unfortunately this check was incorrect and would block until the store
  actually started draining.

This bug is because of channel communication mismatch

### backtrace

```
goroutine 6 [chan receive]:
command-line-arguments.(*consistencyQueue).process(...)
	/root/gobench/gobench/goker/blocking/cockroach/25456/cockroach25456_test.go:51
command-line-arguments.TestCockroach25456(0xc00008e120)
	/root/gobench/gobench/goker/blocking/cockroach/25456/cockroach25456_test.go:77 +0x16c
testing.tRunner(0xc00008e120, 0x550718)
	/usr/local/go/src/testing/testing.go:1050 +0xdc
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1095 +0x28b
```