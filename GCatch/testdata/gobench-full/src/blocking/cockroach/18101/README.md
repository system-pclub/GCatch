# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#18101]|[pull request]|[patch]| Blocking | Resource Deadlock | Double Locking |

[cockroach#18101]:(cockroach18101_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/18101/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/18101
 
## Description


This is some description from previous researchers

> context.Done() signal only stops the goroutine who pulls data
> from a channel, while does not stops goroutines which send data
> to the channel. This causes all goroutines trying to send data
> through the channel to block.

Possible intervening

```
///
/// G1					G2					helper goroutine
/// restore()
/// 					splitAndScatter()
/// <-readyForImportCh
/// 					readyForImportCh<-
/// ...					...
/// 										cancel()
/// return
/// 					readyForImportCh<-
/// -----------------------G2 leak-------------------------
```

### backtrace

```
goroutine 33 [chan send]:
command-line-arguments.splitAndScatter(0x576500, 0xc000078600, 0xc000180000)
    /root/gobench/gobench/goker/blocking/cockroach/18101/cockroach18101_test.go:28 +0x4b
command-line-arguments.restore.func1(0xc000180000, 0x576500, 0xc000078600)
    /root/gobench/gobench/goker/blocking/cockroach/18101/cockroach18101_test.go:15 +0x62
created by command-line-arguments.restore
    /root/gobench/gobench/goker/blocking/cockroach/18101/cockroach18101_test.go:13 +0x70
```