
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#30408]|[pull request]|[patch]| Blocking | Communication Deadlock | Condition Variable |

[moby#30408]:(moby30408_test.go)
[patch]:https://github.com/moby/moby/pull/30408/files
[pull request]:https://github.com/moby/moby/pull/30408
 
## Description

`Wait()` at line 22 has no corresponding `Signal()`.

