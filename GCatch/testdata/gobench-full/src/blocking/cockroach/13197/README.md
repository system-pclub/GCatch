# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#13197]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel & Context |

[cockroach#13197]:(cockroach13197_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/13197/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/13197
 
## Description


This is some description from previous researchers

> One goroutine executing (*Tx).awaitDone() blocks and
> waiting for a signal context.Done().

Possible intervening

```
/// G1 				G2
/// begin()
/// 				awaitDone()
/// 				<-tx.ctx.Done()
/// return
/// -----------G2 leak-------------
```

### backtrace

```
goroutine 19 [chan receive]:
command-line-arguments.(*Tx).awaitDone(0xc000130040)
    /root/gobench/gobench/goker/blocking/cockroach/13197/cockroach13197_test.go:27 +0x4b
created by command-line-arguments.(*DB).begin
    /root/gobench/gobench/goker/blocking/cockroach/13197/cockroach13197_test.go:17 +0xba
```