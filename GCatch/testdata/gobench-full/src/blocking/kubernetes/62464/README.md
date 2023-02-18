
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#62464]|[pull request]|[patch]| Blocking | Resource Deadlock | RWR deadlock |

[kubernetes#62464]:(kubernetes62464_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/62464/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/62464
 
## Description

Some description from developers or pervious reseachers

> This is another example for recursive read lock bug. It has
  been noticed by the go developers that RLock should not be
  recursively used in the same thread.

Possible intervening

```
///
/// G1 									G2
/// m.reconcileState()
/// m.state.GetCPUSetOrDefault()
/// s.RLock()
/// s.GetCPUSet()
/// 									p.RemoveContainer()
/// 									s.GetDefaultCPUSet()
/// 									s.SetDefaultCPUSet()
/// 									s.Lock()
/// s.RLock()
/// ---------------------G1,G2 deadlock---------------------
///
```

