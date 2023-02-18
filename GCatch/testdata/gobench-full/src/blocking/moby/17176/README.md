
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#17176]|[pull request]|[patch]| Blocking | Resource Deadlock | Double locking |

[moby#17176]:(moby17176_test.go)
[patch]:https://github.com/moby/moby/pull/17176/files
[pull request]:https://github.com/moby/moby/pull/17176
 
## Description

Some description from developers or pervious reseachers

> devices.nrDeletedDevices takes devices.Lock() but does
  not drop it if there are no deleted devices. This will block
  other goroutines trying to acquire devices.Lock().
>
> In general reason is that when device deletion is happning,
  we can try deletion/deactivation in a loop. And that that time
  we don't want to block rest of the device operations in parallel.
  So we drop the inner devices lock while continue to hold per
  device lock

Line 36 missing unlock.

