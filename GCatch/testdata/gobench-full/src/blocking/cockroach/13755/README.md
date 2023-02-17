# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#13755]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel & Context |

[cockroach#13755]:(cockroach13755_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/13755/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/13755
 
## Description


This is some description from previous researchers

> The buggy code does not close the db query result (rows),
> so that one goroutine running (*Rows).awaitDone is blocked forever.
> The blocking goroutine is waiting for cancel signal from context.

Possible intervening

```
/// G1 						G2
/// initContextClose()
/// 						awaitDone()
/// 						<-tx.ctx.Done()
/// return
/// ---------------G2 leak-----------------
```

### backtrace

```
goroutine 19 [chan receive]:
command-line-arguments.(*Rows).awaitDone(0xc000102028, 0x5766e0, 0xc000108600)
    /root/gobench/gobench/goker/blocking/cockroach/13755/cockroach13755_test.go:19 +0x48
created by command-line-arguments.(*Rows).initContextClose
    /root/gobench/gobench/goker/blocking/cockroach/13755/cockroach13755_test.go:15 +0x82
```