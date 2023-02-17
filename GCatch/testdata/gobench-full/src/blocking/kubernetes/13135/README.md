
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[kubernetes#13135]|[pull request]|[patch]| Blocking | Resource Deadlock | AB-BA deadlock |

[kubernetes#13135]:(kubernetes13135_test.go)
[patch]:https://github.com/kubernetes/kubernetes/pull/13135/files
[pull request]:https://github.com/kubernetes/kubernetes/pull/13135
 
## Description

```
///
/// G1								G2								G3
/// NewCacher()
/// watchCache.SetOnReplace()
/// watchCache.SetOnEvent()
/// 								cacher.startCaching()
///									c.Lock()
/// 								c.reflector.ListAndWatch()
/// 								r.syncWith()
/// 								r.store.Replace()
/// 								w.Lock()
/// 								w.onReplace()
/// 								cacher.initOnce.Do()
/// 								cacher.Unlock()
/// return cacher
///																	c.watchCache.Add()
///																	w.processEvent()
///																	w.Lock()
///									cacher.startCaching()
///									c.Lock()
///									...
///																	c.Lock()
///									w.Lock()
///--------------------------------G2,G3 deadlock-------------------------------------
///
```

