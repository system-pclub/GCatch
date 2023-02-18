
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#4395]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[moby#4395]:(moby4395_test.go)
[patch]:https://github.com/moby/moby/pull/4395/files
[pull request]:https://github.com/moby/moby/pull/4395
 
## Description

Some description from developers or pervious reseachers

> The anonyous goroutine could be waiting on sending to
  the channel which might never be drained.

Possible intervening

```
///
/// G1				G2
/// Go()
/// return ch
/// 				ch <- f()
/// ----------G2 leak-------------
///
```

## Backtrace

```
```

