
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#30872]|[pull request]|[patch]| Blocking | Resource Deadlock | AB-BA deadlock |

[kubernetes#30872]:(kubernetes30872_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/30872/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/30872
 
## Description

This is a AB-BA deadlock. (Lock acquires at line 92 and at line 157 respectively)

## Backtrace

```
goroutine 7 [semacquire]:
sync.runtime_SemacquireMutex(0xc00000c0a4, 0xc000072200, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc00000c0a0)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
command-line-arguments.(*federatedInformerImpl).addCluster(0xc00000c0a0)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:93 +0x78
command-line-arguments.NewFederatedInformer.func1()
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:173 +0x2a
command-line-arguments.ResourceEventHandlerFuncs.OnAdd(0xc00005e490)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:71 +0x33
command-line-arguments.NewInformer.func1()
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:184 +0x2f
command-line-arguments.(*DeltaFIFO).Pop(0xc0000180a0, 0xc00000c0c0)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:165 +0x5f
command-line-arguments.(*Controller).processLoop(...)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:53
command-line-arguments.JitterUntil.func1(...)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:25
command-line-arguments.JitterUntil(0xc000100fb0, 0xc000076300)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:26 +0x2a
command-line-arguments.Util(...)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:14
command-line-arguments.(*Controller).Run(0xc00000c0e0, 0xc000076300)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:45 +0x53
created by command-line-arguments.(*federatedInformerImpl).Start
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:102 +0xb8

 Goroutine 8 in state semacquire, with sync.runtime_SemacquireMutex on top of the stack:
goroutine 8 [semacquire]:
sync.runtime_SemacquireMutex(0xc00000c0a4, 0x0, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc00000c0a0)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
command-line-arguments.(*federatedInformerImpl).Stop(0xc00000c0a0)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:106 +0x8e
command-line-arguments.(*NamespaceController).Run.func1(0xc0000762a0, 0xc00000c080)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:146 +0x4b
created by command-line-arguments.(*NamespaceController).Run
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:144 +0x64

 Goroutine 9 in state semacquire, with sync.runtime_SemacquireMutex on top of the stack:
goroutine 9 [semacquire]:
sync.runtime_SemacquireMutex(0xc0000180a4, 0x1, 0x1)
    /usr/local/go/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc0000180a0)
    /usr/local/go/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(...)
    /usr/local/go/src/sync/mutex.go:81
sync.(*RWMutex).Lock(0xc0000180a0)
    /usr/local/go/src/sync/rwmutex.go:98 +0x97
command-line-arguments.(*DeltaFIFO).HasSynced(0xc0000180a0)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:158 +0x3a
command-line-arguments.(*Controller).HasSynced(0xc00000c0e0)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:49 +0x33
command-line-arguments.(*federatedInformerImpl).ClustersSynced(0xc00000c0a0)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:89 +0x6d
command-line-arguments.(*NamespaceController).isSynced(...)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:135
command-line-arguments.(*NamespaceController).reconcileNamespace(...)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:139
command-line-arguments.(*NamespaceController).Run.func2()
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:149 +0x34
command-line-arguments.(*DelayingDeliverer).StartWithHandler.func1(0xc00005e4a0)
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:115 +0x25
created by command-line-arguments.(*DelayingDeliverer).StartWithHandler
    /root/gobench/goker/blocking/kubernetes/30872/kubernetes30872_test.go:114 +0x3f
```

