
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#5509]|[pull request]|[patch]| Blocking | Resource Deadlock | Double locking |

[etcd#5509]:(etcd5509_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/5509/files
[pull request]:https://github.com/etcd-io/etcd/pull/5509
 
## Description

Some description from developers or pervious reseachers

> r.acquire() returns holding r.client.mu.RLock() on success; 
> it was dead locking because it was returning with the rlock held on 
> a failure path and leaking it. After that any call to client.Close() 
> will block forever waiting for the wlock.

Line 42 : Missing RUnlock before return 


