
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#11298]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel & Condition Variable |

[kubernetes#11298]:(kubernetes11298_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/11298/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/11298
 
## Description

Some description from developers or pervious reseachers

> n.node used the n.lock as underlaying locker. The service loop initially
  locked it, the Notify function tried to lock it before calling n.node.Signal,
  leading to a dead-lock.

`n.cond.Signal()` at line 59 and line 81 are not guaranteed to `n.cond.Wait` at line 56.

