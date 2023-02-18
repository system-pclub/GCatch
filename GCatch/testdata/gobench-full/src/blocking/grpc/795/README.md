
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#795]|[pull request]|[patch]| Blocking | Resource Deadlock | Double locking |

[grpc#795]:(grpc795_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/795/files
[pull request]:https://github.com/grpc/grpc-go/pull/795
 
## Description

line 20 missing unlock
