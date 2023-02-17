
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[istio#8144]|[pull request]|[patch]| NonBlocking | Traditional | Data race |

[istio#8144]:(istio8144_test.go)
[patch]:https://github.com/istio/istio/pull/8144/files
[pull request]:https://github.com/istio/istio/pull/8144
 

## Backtrace

```
Write at 0x00c0000ac088 by goroutine 9:
  command-line-arguments.(*callbackRecorder).callback()
      /root/gobench/goker/nonblocking/istio/8144/istio8144_test.go:15 +0x5e
  command-line-arguments.(*callbackRecorder).callback-fm()
      /root/gobench/goker/nonblocking/istio/8144/istio8144_test.go:14 +0x22
  command-line-arguments.(*ttlCache).evictExpired.func1()
      /root/gobench/goker/nonblocking/istio/8144/istio8144_test.go:29 +0x5a
  sync.(*Map).Range()
      /usr/local/go/src/sync/map.go:333 +0x155
  command-line-arguments.(*ttlCache).evictExpired()
      /root/gobench/goker/nonblocking/istio/8144/istio8144_test.go:28 +0x5d
  command-line-arguments.(*ttlCache).evicter()
      /root/gobench/goker/nonblocking/istio/8144/istio8144_test.go:24 +0x38

Previous read at 0x00c0000ac088 by goroutine 8:
  command-line-arguments.TestIstio8144.func1()
      /root/gobench/goker/nonblocking/istio/8144/istio8144_test.go:54 +0x15d

Goroutine 9 (running) created at:
  command-line-arguments.NewTTLWithCallback()
      /root/gobench/goker/nonblocking/istio/8144/istio8144_test.go:42 +0xc0
  command-line-arguments.TestIstio8144.func1()
      /root/gobench/goker/nonblocking/istio/8144/istio8144_test.go:52 +0x115

Goroutine 8 (finished) created at:
  command-line-arguments.TestIstio8144()
      /root/gobench/goker/nonblocking/istio/8144/istio8144_test.go:49 +0xa2
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1050 +0x1eb
```

