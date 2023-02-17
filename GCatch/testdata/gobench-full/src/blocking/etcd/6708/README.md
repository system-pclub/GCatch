
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[etcd#6708]|[pull request]|[patch]| Blocking | Resource Deadlock | Double locking |

[etcd#6708]:(etcd6708_test.go)
[patch]:https://github.com/etcd-io/etcd/pull/6708/files
[pull request]:https://github.com/etcd-io/etcd/pull/6708
 
## Description

Line 54, 49 double locking


