
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#7559]|[pull request]|[patch]| Blocking | Resource Deadlock | Double locking |

[moby#7559]:(moby7559_test.go)
[patch]:https://github.com/moby/moby/pull/7559/files
[pull request]:https://github.com/moby/moby/pull/7559
 
## Description

Line 25 missing unlock

