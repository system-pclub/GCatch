
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#4951]|[pull request]|[patch]| Blocking | Resource Deadlock | AB-BA deadlock |

[moby#4951]:(moby4951_test.go)
[patch]:https://github.com/moby/moby/pull/4951/files
[pull request]:https://github.com/moby/moby/pull/4951
 
## Description

Some description from developers or pervious reseachers

> The root cause and patch is clearly explained in the commit
  description. The global lock is devices.Lock(), and the device
  lock is baseInfo.lock.Lock(). It is very likely that this bug
  can be reproduced.


