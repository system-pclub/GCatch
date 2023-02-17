
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[syncthing#4829]|[pull request]|[patch]| Blocking | Resource Deadlock | Double locking |

[syncthing#4829]:(syncthing4829_test.go)
[patch]:https://github.com/syncthing/syncthing/pull/4829/files
[pull request]:https://github.com/syncthing/syncthing/pull/4829
 
## Description

Double locking at line 17 and line 30

