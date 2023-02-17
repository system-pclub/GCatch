
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#29733]|[pull request]|[patch]| Blocking | Communication Deadlock | Condition Variable |

[moby#29733]:(moby29733_test.go)
[patch]:https://github.com/moby/moby/pull/29733/files
[pull request]:https://github.com/moby/moby/pull/29733
 
## Description

`Wait()` at line 21 has no corresponding `Signal()`.

