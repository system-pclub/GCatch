
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[istio#16742]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[istio#16742]:(istio16742_test.go)
[patch]:https://github.com/istio/istio/pull/16742/files
[pull request]:https://github.com/istio/istio/pull/16742
 

## Backtrace

```
Write at 0x00c00000e080 by goroutine 10:
  command-line-arguments.(*DiscoveryServer).WorkloadUpdate()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:72 +0x11a
  command-line-arguments.(*MemServiceDiscovery).AddWorkload()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:86 +0x85
  command-line-arguments.TestIstio16742.func1.2()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:105 +0x66

Previous read at 0x00c00000e080 by goroutine 9:
  command-line-arguments.(*ConfigGeneratorImpl).buildSidecarOutboundHTTPRouteConfig()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:28 +0x3b
  command-line-arguments.(*ConfigGeneratorImpl).BuildHTTPRoutes()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:24 +0x32
  command-line-arguments.(*DiscoveryServer).generateRawRoutes()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:62 +0x11e
  command-line-arguments.(*DiscoveryServer).pushRoute()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:66 +0xd0
  command-line-arguments.(*DiscoveryServer).StreamAggregatedResources()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:58 +0xcf
  command-line-arguments.TestIstio16742.func1.1()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:101 +0x95

Goroutine 10 (running) created at:
  command-line-arguments.TestIstio16742.func1()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:103 +0x15d

Goroutine 9 (finished) created at:
  command-line-arguments.TestIstio16742.func1()
      /root/gobench/goker/nonblocking/istio/16742/istio16742_test.go:99 +0x12e
```

