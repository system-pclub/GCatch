
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#33781]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel & Context |

[moby#33781]:(moby33781_test.go)
[patch]:https://github.com/moby/moby/pull/33781/files
[pull request]:https://github.com/moby/moby/pull/33781
 
## Description

Some description from developers or pervious reseachers

> The goroutine created using anonymous function is blocked at
  sending message to a unbuffered channel. However there exists a
  path in the parent goroutine where the parent function will
  return without draining the channel.

Possible intervening

```
///
/// G1 				G2				G3
/// monitor()
/// <-time.After()
/// 				stop <-
/// <-stop
/// 				return
/// cancelProbe()
/// return
/// 								result<-
///----------------G3 leak------------------
///
```

