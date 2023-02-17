
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[syncthing#5795]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[syncthing#5795]:(syncthing5795_test.go)
[patch]:https://github.com/syncthing/syncthing/pull/5795/files
[pull request]:https://github.com/syncthing/syncthing/pull/5795
 
## Description

`<-c.dispatcherLoopStopped` at line 82 is blocking forever because
dispatcherLoop() is blocking at line 72.

