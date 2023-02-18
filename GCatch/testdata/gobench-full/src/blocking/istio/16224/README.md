
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[istio#16224]|[pull request]|[patch]| Blocking | Mixed Deadlock | Channel & Lock |

[istio#16224]:(istio16224_test.go)
[patch]:https://github.com/istio/istio/pull/16224/files
[pull request]:https://github.com/istio/istio/pull/16224
 
## Description

A goroutine holds a mutex at line 91 and then blocked at line 93.
Another goroutine attempt acquire the same mutex at line 101 to 
further drains the same channel at 103.

