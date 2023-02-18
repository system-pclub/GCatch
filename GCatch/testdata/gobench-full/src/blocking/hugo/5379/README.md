
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[hugo#5379]|[pull request]|[patch]| Blocking | Resource Deadlock | Double locking |

[hugo#5379]:(hugo5379_test.go)
[patch]:https://github.com/gohugoio/hugo/pull/5379/files
[pull request]:https://github.com/gohugoio/hugo/pull/5379
 
## Description

A goroutine first acquire `contentInitMu` at line 99 then
acquire the same Mutex at line 66
