# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#35073]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[cockroach#35073]:(cockroach35073_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/35073/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/35073
 
## Description


This is some description from developers

> Previously, the outbox could fail during startup without closing its
  RowChannel. This could lead to deadlocked flows in rare cases.

Channel communication mismatch

### backtrace

```
goroutine 18 [chan send]:
command-line-arguments.(*RowChannel).Push(...)
	/root/gobench/gobench/goker/blocking/cockroach/35073/cockroach35073_test.go:49
command-line-arguments.TestCockroach35073(0xc000144120)
	/root/gobench/gobench/goker/blocking/cockroach/35073/cockroach35073_test.go:110 +0x222
testing.tRunner(0xc000144120, 0x5519d8)
	/usr/local/go/src/testing/testing.go:1050 +0xdc
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1095 +0x28b

goroutine 19 [chan send]:
command-line-arguments.(*RowChannel).Push(...)
	/root/gobench/gobench/goker/blocking/cockroach/35073/cockroach35073_test.go:49
command-line-arguments.TestCockroach35073.func1(0xc000104480, 0xc0001180a0)
	/root/gobench/gobench/goker/blocking/cockroach/35073/cockroach35073_test.go:103 +0x62
created by command-line-arguments.TestCockroach35073
	/root/gobench/gobench/goker/blocking/cockroach/35073/cockroach35073_test.go:102 +0x18c
```