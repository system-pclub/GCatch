
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#33293]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[moby#33293]:(moby33293_test.go)
[patch]:https://github.com/moby/moby/pull/33293/files
[pull request]:https://github.com/moby/moby/pull/33293
 
## Description

Possible intervening

```
///
/// G1
/// containerWait()
/// errC <- err
/// ---------G1 leak---------------
///
```
