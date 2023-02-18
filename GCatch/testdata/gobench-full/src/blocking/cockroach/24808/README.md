# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#24808]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[cockroach#24808]:(cockroach24808_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/24808/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/24808
 
## Description


This is some description from developers

> When we Start the Compactor, it may already have received
> Suggestions, deadlocking the previously blocking write to a full
> channel.

### backtrace

```
goroutine 33 [chan send]:
command-line-arguments.(*Compactor).Start(0xc0000ce010, 0x574b00, 0xc0000aa010, 0xc0000b4040)
	/root/gobench/gobench/goker/blocking/cockroach/24808/cockroach24808_test.go:50 +0x3c
command-line-arguments.TestCockroach24808(0xc0000d0120)
	/root/gobench/gobench/goker/blocking/cockroach/24808/cockroach24808_test.go:70 +0x18f
testing.tRunner(0xc0000d0120, 0x552e90)
	/usr/local/go/src/testing/testing.go:1050 +0xdc
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1095 +0x28b
```