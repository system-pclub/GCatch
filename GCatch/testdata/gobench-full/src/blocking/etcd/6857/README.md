
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#6857]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[etcd#6857]:(etcd6857_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/6857/files
[pull request]:https://github.com/etcd-io/etcd/pull/6857
 
## Description

Possible intervening

```
///
/// G1				G2				G3
/// n.run()
///									n.Stop()
///									n.stop<-
/// <-n.stop
///									<-n.done
/// close(n.done)
///	return
///									return
///					n.Status()
///					n.status<-
///----------------G2 leak-------------------
///
```

