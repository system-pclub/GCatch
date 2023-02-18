# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[cockroach#2448]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[cockroach#2448]:(cockroach2448_test.go)
[patch]:https://github.com/cockroachdb/cockroach/pull/2448/files
[pull request]:https://github.com/cockroachdb/cockroach/pull/2448
 
## Description

This is some description from previous researchers

>  This bug is caused by two goroutines waiting for each other
>  to unblock their channels.
>
>  (1) MultiRaft sends the commit event for the Membership change
>
>  (2) store.processRaft takes it and begins processing
>
>  (3) another command commits and triggers another sendEvent, but
  	   this blocks since store.processRaft isn't ready for another
  	   select. Consequently the main MultiRaft loop is waiting for
  	   that as well.
>
>  (4) the Membership change was applied to the range, and the store
       now tries to execute the callback
>
>  (5) the callback tries to write to callbackChan, but that is
  	   consumed by the MultiRaft loop, which is currently waiting
  	   for store.processRaft to consume from the events channel,
  	   which it will only do after the callback has completed.

Possible intervening
```
G1								G2
s.processRaft()
e := <-s.multiraft.Events
								st.start()
 								s.handleWriteResponse()
 								s.processCommittedEntry()
 								s.sendEvent()
 								m.Events <- event
								...
 								s.handleWriteResponse()
 								s.processCommittedEntry()
 								s.sendEvent()
 								m.Events <- event
callback()
s.callbackChan <- func()
```

### backtrace

```
goroutine 19 [select]:
command-line-arguments.(*state).processCommittedEntry.func1()
    /root/gobench/gobench/goker/blocking/cockroach/2448/cockroach2448_test.go:57 +0xbc
command-line-arguments.(*Store).processRaft(0xc0000ae038)
    /root/gobench/gobench/goker/blocking/cockroach/2448/cockroach2448_test.go:80 +0xfc
created by command-line-arguments.TestCockroach2448
    /root/gobench/gobench/goker/blocking/cockroach/2448/cockroach2448_test.go:106 +0x1d7

 Goroutine 20 in state select, with command-line-arguments.(*MultiRaft).sendEvent on top of the stack:
goroutine 20 [select]:
command-line-arguments.(*MultiRaft).sendEvent(0xc0000be040, 0x50fc80, 0xc00000e010)
    /root/gobench/gobench/goker/blocking/cockroach/2448/cockroach2448_test.go:28 +0xc6
command-line-arguments.(*state).processCommittedEntry(0xc0000ae030)
    /root/gobench/gobench/goker/blocking/cockroach/2448/cockroach2448_test.go:55 +0x91
command-line-arguments.(*state).handleWriteResponse(...)
    /root/gobench/gobench/goker/blocking/cockroach/2448/cockroach2448_test.go:51
command-line-arguments.(*state).start(0xc0000ae030)
    /root/gobench/gobench/goker/blocking/cockroach/2448/cockroach2448_test.go:46 +0xfe
created by command-line-arguments.TestCockroach2448
    /root/gobench/gobench/goker/blocking/cockroach/2448/cockroach2448_test.go:107 +0x1f9
```