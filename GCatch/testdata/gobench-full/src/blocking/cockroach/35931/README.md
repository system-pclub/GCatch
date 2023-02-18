# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#35931]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[cockroach#35931]:(cockroach35931_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/35931/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/35931
 
## Description


This is some description from developers

> Previously, if a processor that reads from multiple inputs was waiting
  on one input to provide more data, and the other input was full, and
  both inputs were connected to inbound streams, it was possible to
  deadlock the system during flow cancellation when trying to propagate
  the cancellation metadata messages into the flow. The problem was that
  the cancellation method wrote metadata messages to each inbound stream
  one at a time, so if the first one was full, the canceller would block
  and never send a cancellation message to the second stream, which was
  the one actually being read from.

Channel mismatch

### backtrace

```
goroutine 6 [chan send]:
command-line-arguments.(*RowChannel).Push(0xc00000e038)
	/root/gobench/gobench/goker/blocking/cockroach/35931/cockroach35931_test.go:22 +0x38
command-line-arguments.(*Flow).cancel(0xc00000c080)
	/root/gobench/gobench/goker/blocking/cockroach/35931/cockroach35931_test.go:69 +0xa9
command-line-arguments.TestCockroach35931(0xc00010e120)
	/root/gobench/gobench/goker/blocking/cockroach/35931/cockroach35931_test.go:113 +0x322
testing.tRunner(0xc00010e120, 0x552258)
	/usr/local/go/src/testing/testing.go:1050 +0xdc
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1095 +0x28b
```