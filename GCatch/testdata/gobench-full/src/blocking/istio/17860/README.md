
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[istio#17860]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[istio#17860]:(istio17860_test.go)
[patch]:https://github.com/istio/istio/pull/17860/files
[pull request]:https://github.com/istio/istio/pull/17860
 
## Description

`a.statusCh` can't be drained at line 70.

## Backtace

```
goroutine 33 [chan send]:
command-line-arguments.(*agent).runWait(0xc000078300, 0x2)
    /root/gobench/goker/blocking/istio/17860/istio17860_test.go:71 +0x43
created by command-line-arguments.(*agent).Restart
    /root/gobench/goker/blocking/istio/17860/istio17860_test.go:67 +0xd0

 Goroutine 7 in state chan send, with command-line-arguments.(*agent).runWait on top of the stack:
goroutine 7 [chan send]:
command-line-arguments.(*agent).runWait(0xc000078300, 0x1)
    /root/gobench/goker/blocking/istio/17860/istio17860_test.go:71 +0x43
created by command-line-arguments.(*agent).Restart
    /root/gobench/goker/blocking/istio/17860/istio17860_test.go:67 +0xd0
]
```

