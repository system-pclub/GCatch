
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#26980]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[kubernetes#26980]:(kubernetes26980_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/26980/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/26980
 
## Description

A goroutine holds a mutex at line 24 and blocked at line 35.
Another goroutine blocked at line 58 by acquiring the same mutex. 

