
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#36114]|[pull request]|[patch]| Blocking | Resource Deadlock | Double locking |

[moby#36114]:(moby36114_test.go)
[patch]:https://github.com/moby/moby/pull/36114/files
[pull request]:https://github.com/moby/moby/pull/36114
 
## Description

Some description from developers or pervious reseachers

> This is a double lock bug. The the lock for the
  struct svm has already been locked when calling
  svm.hotRemoveVHDsAtStart()

