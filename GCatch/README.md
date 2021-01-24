# What is GCatch and how to use it

1. What is GCatch?

A static checker that takes Golang source code as input and detects concurrency bugs in source code.

2. How to install and run GCatch?

- run "sudo installZ3.sh" to install Z3

- run "install.sh" to install GCatch

- run "run.sh" to run GCatch on testdata/grpc-buggy

3. What are the checkers in GCatch?

  - Forget Unlock
    - When a function returns, if there are any Mutex or RWMutex that are previously locked in the same function but not unlocked, report a bug.
  - Struct Field
    - If a field of a structure is often protected by a Mutex/RWMutex, but there are a few times that it is not protected, it is likely that the programmer forgot to use a mutex. Report a bug and show all usages of this field.
  - Double Lock
    - When a Mutex/RWMutex is locked in one function, and before it is unlocked, some other functions are called and the Mutex/RWMutex is locked again, report a bug.
  - Conflict Lock
    - Consider two Mutex/RWMutex m1 and m2. When one goroutine runs m1.Lock() and m2.Lock(), and another goroutine runs m2.Lock() and then m1.Lock(), report a bug.
  - BMOC
    - When a channel operation blocks forever, report a bug.
  - Fatal
    - If a testing function creates a goroutine that calls testing.Fatal()/FailNow()/Skip()/SkipNow(), report a bug.

# Introduction of each package

1. cmd contains the main() function of GCatch. When you want to install GCatch, you need to set GOPATH to this repo, open cmd/GCatch, and run `go install`

2. analysis contains post-dominator analysis and code to analyze the results of pointer analysis

3. checkers contains the code to run each checker. The main() function will invoke checkers in this package

4. config contains the configure variables, global variables

5. instoinfo contains the definition of synchronization primitives, including channels and mutexes (named locker in GCatch)

6. output contains usaful functions to print to terminal

7. path contains some CFG analysis

8. ssabuild contains the code to build the AST and SSA from the input program

9. syncgraph is the core of BMOC checker. It contains a definition of SyncGraph, a data structure that records all the CFG, callgraph, dependency and alias information we need to detect a BMOC bug. It also contains the code to generate Z3 constraints and invoke Z3.

10. tests contains some functions used to test if traditional checkers work well

11. tools are copied vendor packages like golang.org/x/tools, because we want to maintain our own copies of them.

12. util contains some functions used by all other packages