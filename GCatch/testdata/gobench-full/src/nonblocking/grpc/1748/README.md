
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#1748]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[grpc#1748]:(grpc1748_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/1748/files
[pull request]:https://github.com/grpc/grpc-go/pull/1748
 

## Backtrace

```
Read at 0x000000735418 by goroutine 14:
  command-line-arguments.(*addrConn).resetTransport()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:65 +0x3c
  command-line-arguments.(*addrConn).transportMonitor()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:69 +0x2c
  command-line-arguments.(*addrConn).connect.func1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:74 +0x2b

Previous write at 0x000000735418 by goroutine 8:
  command-line-arguments.TestGrpc1748.func1.1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:143 +0x3a
  command-line-arguments.TestGrpc1748.func1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:147 +0x120

Goroutine 14 (running) created at:
  command-line-arguments.(*addrConn).connect()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:73 +0x4c
  command-line-arguments.(*acBalancerWrapper).Connect()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:81 +0x92
  command-line-arguments.(*pickfirstBalancer).HandleResolvedAddrs()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:46 +0xc4
  command-line-arguments.(*ccBalancerWrapper).watcher()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:111 +0xb4

Goroutine 8 (finished) created at:
  command-line-arguments.TestGrpc1748()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:138 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```


```
Read at 0x000000735418 by goroutine 16:
  command-line-arguments.(*addrConn).resetTransport()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:65 +0x3c
  command-line-arguments.(*addrConn).transportMonitor()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:69 +0x2c
  command-line-arguments.(*addrConn).connect.func1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:74 +0x2b

Previous write at 0x000000735418 by goroutine 8:
  command-line-arguments.TestGrpc1748.func1.1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:143 +0x3a
  command-line-arguments.TestGrpc1748.func1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:147 +0x120

Goroutine 16 (running) created at:
  command-line-arguments.(*addrConn).connect()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:73 +0x4c
  command-line-arguments.(*acBalancerWrapper).Connect()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:81 +0x92
  command-line-arguments.(*pickfirstBalancer).HandleResolvedAddrs()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:46 +0xc4
  command-line-arguments.(*ccBalancerWrapper).watcher()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:111 +0xb4

Goroutine 8 (finished) created at:
  command-line-arguments.TestGrpc1748()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:138 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

```
Read at 0x000000735418 by goroutine 12:
  command-line-arguments.(*addrConn).resetTransport()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:65 +0x3c
  command-line-arguments.(*addrConn).transportMonitor()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:69 +0x2c
  command-line-arguments.(*addrConn).connect.func1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:74 +0x2b

Previous write at 0x000000735418 by goroutine 8:
  command-line-arguments.TestGrpc1748.func1.1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:143 +0x3a
  command-line-arguments.TestGrpc1748.func1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:147 +0x120

Goroutine 12 (running) created at:
  command-line-arguments.(*addrConn).connect()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:73 +0x4c
  command-line-arguments.(*acBalancerWrapper).Connect()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:81 +0x92
  command-line-arguments.(*pickfirstBalancer).HandleResolvedAddrs()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:46 +0xc4
  command-line-arguments.(*ccBalancerWrapper).watcher()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:111 +0xb4

Goroutine 8 (finished) created at:
  command-line-arguments.TestGrpc1748()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:138 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

```
Read at 0x000000735418 by goroutine 11:
  command-line-arguments.(*addrConn).resetTransport()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:65 +0x3c
  command-line-arguments.(*addrConn).transportMonitor()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:69 +0x2c
  command-line-arguments.(*addrConn).connect.func1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:74 +0x2b

Previous write at 0x000000735418 by goroutine 8:
  command-line-arguments.TestGrpc1748.func1.1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:143 +0x3a
  command-line-arguments.TestGrpc1748.func1()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:147 +0x120

Goroutine 11 (running) created at:
  command-line-arguments.(*addrConn).connect()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:73 +0x4c
  command-line-arguments.(*acBalancerWrapper).Connect()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:81 +0x92
  command-line-arguments.(*pickfirstBalancer).HandleResolvedAddrs()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:46 +0xc4
  command-line-arguments.(*ccBalancerWrapper).watcher()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:111 +0xb4

Goroutine 8 (finished) created at:
  command-line-arguments.TestGrpc1748()
      /root/gobench/goker/nonblocking/grpc/1748/grpc1748_test.go:138 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```
