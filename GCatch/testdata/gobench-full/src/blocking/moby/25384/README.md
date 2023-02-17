
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[moby#25384]|[pull request]|[patch]| Blocking | Mixed Deadlock | Misuse WaitGroup |

[moby#25384]:(moby25384_test.go)
[patch]:https://github.com/moby/moby/pull/25384/files
[pull request]:https://github.com/moby/moby/pull/25384
 
## Description

Some description from developers or pervious reseachers

> When n=1 (len(pm.plugins)), the location of group.Wait() doesn’t matter.
  When n is larger than 1, group.Wait() is invoked in each iteration. Whenever
  group.Wait() is invoked, it waits for group.Done() to be executed n times.
  However, group.Done() is only executed once in one iteration.

Misuse of sync.WaitGroup

