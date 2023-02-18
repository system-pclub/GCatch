
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#70277]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[kubernetes#70277]:kubernetes70277_test.go
[patch]:https://github.com/kubernetes/kubernetes/pull/70277/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/70277
 
## Description

Some description from developers or pervious reseachers

> wait.poller() returns a function with type WaitFunc. 
> the function creates a goroutine and the goroutine only 
> quits when after or done closed.


The doneCh defined at line 70 is never closed.
