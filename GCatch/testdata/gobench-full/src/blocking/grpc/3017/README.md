
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#3017]|[pull request]|[patch]| Blocking | Resource Deadlock | Double locking |

[grpc#3017]:(grpc3017_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/3017/files
[pull request]:https://github.com/grpc/grpc-go/pull/3017
 
## Description

Line 65 missing unlock

