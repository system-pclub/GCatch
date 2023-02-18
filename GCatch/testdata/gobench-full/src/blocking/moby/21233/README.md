
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#21233]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[moby#21233]:(moby21233_test.go)
[patch]:https://github.com/moby/moby/pull/21233/files
[pull request]:https://github.com/moby/moby/pull/21233
 
## Description

Some description from developers or pervious reseachers

> This test was checking that it received every progress update that was
  produced. But delivery of these intermediate progress updates is not
  guaranteed. A new update can overwrite the previous one if the previous
  one hasn't been sent to the channel yet.
  
> The call to t.Fatalf exited the cur rent goroutine which was consuming
  the channel, which caused a deadlock and eventual test timeout rather
  than a proper failure message.

Possible intervening

```
///
/// G1 						G2					G3
/// testTransfer()
/// tm.Transfer()
/// t.Watch()
/// 						WriteProgress()
/// 						ProgressChan<-
/// 											<-progressChan
/// 						...					...
/// 						return
/// 											<-progressChan
/// <-watcher.running
/// ----------------------G1, G3 leak--------------------------
///
```

