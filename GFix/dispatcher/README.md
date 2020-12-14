# GoConcurrencyChecker

1. What is GoConcurrencyChecker?

A static checker that takes Golang source code as input and points out concurrency bugs in source code.  The tool is still under development. 

2. How does it work?

(1). Overall procedure
  -  Parses the source code and builds SSA of the whole program on it (How to build SSA on source code). Ignores code in "vendor".
  -  Runs checkers [A1] ~ [A8], and reports bugs if any.
  -  Builds the call-graph that covers functions used in *_test.go files.
  -  Run checker [B1], which depends on call-graph, and reports bugs if any.

(2). Introduction to checkers
(Note: some details like how to avoid false positives are not presented)
  - Checker C4: Missing-Unlock
    - When a function returns, if there are any Mutex or RWMutex that are previously locked in the same function but not unlocked, report a bug.
  - Checker C3: Inconsistent-Field-Protection
    - If a field of a structure is always protected by at least one Mutex/RWMutex, but there are a few times that it is not protected, it is likely that the programmer forgot to use a mutex. Report a bug and show all usages of this field.
  - Checker C5: Double-Lock
    - When a Mutex/RWMutex is locked in one function, and before it is unlocked, some other functions are called and the Mutex/RWMutex is locked again, report a bug.
  - Checker C6: Channel-In-Critical-Section
    - If there are two goroutines, one sending to a channel that is protected by a Mutex/RWMutex, and the other receiving from the same channel that is also protected by the same Mutex/RWMutex, report a bug.
  - Checker C7: Goroutine-Leak
    - If an unbuffered channel is only used by two goroutines, one guaranteeing using the channel, and the other only using the channel in some select cases, report a bug.
  - Checker C8: API-Fatal
    - If a testing function creates a goroutine that uses testing.Fatal()/FailNow()/Skip()/SkipNow(), report a bug.
  - Checker C1: API-Context
    - If context.Context is used by functions like fmt.Sprintf(), report a bug.
  - Checker C2: API-Atomic64
    - Give warning about usage of 64-bit functions in atomic package, which will cause error in old machines.
  - Checker B1: Anonymous-Function-Race
    - When a parent function creates an anonymous function, if any variables used in the anonymous function are changed later in the parent function, report a bug.

3. What has it found?
Up to now, our checker has found:
- 6 bugs in Docker, 6 of which are verified by developers;
- 49 bugs in etcd, 17 of which are verified by developers;
- 7 bugs in grpc, 7 of which is verified by developers;
- 20 bugs in Kubernetes, 13 of which are verified by developers;
- 10 bugs in cockroachdb, 6 of which are verified by developers;
- 6 bugs in bboltdb;
- 4 bugs in other software like frp, Syncthing, golang.com/x/time, ..., 3 of which are verified by developers.

4. How to use it?

- Open terminal, and use the following commands:
  - mkdir newdir
  - cd newdir
  - export GOPATH=`pwd`
  - mkdir src bin pkg
  - cd src
  - mkdir git.gradebot.org
  - cd git.gradebot.org
  - mkdir zxl381
  - cd zxl381
  - git clone git@git.gradebot.org:zxl381/goconcurrencychecker.git
  - cd ./goconcurrencychecker/cmd/staticchecker
  - go install
  - cd $GOPATH/bin
  - export GOPATH=/GOPATH/of/the/project/to/be/scanned
  - ./staticchecker $PARAMETERS

- Required $PARAMETERS:
  - -path=...  
    - Full path of the project you want to scan. This should be the path of the directory that contains all packages in your project.
  - -include=... 
    - Relative path (what's after /src/) of the project you want to scan. This parameter may seem to be unnecessary, but we need it to avoid some errors.

- Commonly used $PARAMETERS:
  - -compile-error
    - If the project has compile errors, print these compilation errors
  - -exclude=...
    - Name of directories that you want to ignore, divided by ":"
    - Default value: "vendor"
  - -thread=...
    - Number of threads you can assign to the checker. 
    - Default value: 2
  - -race
    - Turn on the "Race in anonymous function" checker. This checker is expensive but powerful

Example:
I want to check bugs in docker. I will do the following:
- cd ${one_directory}
- export GOPATH=/User/me/Desktop/docker
- ./staticchecker -path=/Users/me/Desktop/docker/src/github.com/docker/docker -include=github.com/docker/docker -race

