# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#10790]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel & Context |

[cockroach#10790]:(cockroach10790_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/10790/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/10790
 
## Description

This is some description from previous researchers

> It is possible that a message from ctxDone will make the function beginCmds
> returns without draining the channel ch, so that goroutines created by anonymous
> function will leak.

Possible intervening

```
///
/// G1					G2				helper goroutine
/// 									r.sendChans()
/// r.beginCmds()
/// 									ch1 <- true
/// <- ch1
///										ch2 <- true
///	...					...				...
///						cancel()
///	<- ch1
///	------------------G1 leak--------------------------
```

### backtrace

```
goroutine 603 [chan receive]:
command-line-arguments.(*Replica).beginCmds.func1(0xc00000c5a0)
    /root/gobench/gobench/goker/blocking/cockroach/10790/cockroach10790_test.go:51 +0x52
created by command-line-arguments.(*Replica).beginCmds
    /root/gobench/gobench/goker/blocking/cockroach/10790/cockroach10790_test.go:49 +0x13f

 Goroutine 604 in state chan receive, with command-line-arguments.(*Replica).beginCmds.func1 on top of the stack:
goroutine 604 [chan receive]:
command-line-arguments.(*Replica).beginCmds.func1(0xc00000c5a0)
    /root/gobench/gobench/goker/blocking/cockroach/10790/cockroach10790_test.go:51 +0x52
created by command-line-arguments.(*Replica).beginCmds
    /root/gobench/gobench/goker/blocking/cockroach/10790/cockroach10790_test.go:49 +0x13f
```