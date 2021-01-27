# GCatch: Automatically Detecting Concurrency Bugs in Go

## Descriptions

GCatch contains a suite of static detectors aiming to identify concurrency bugs in large, real Go software systems. Concurrency bugs covered by the current version of GCatch include blocking misuse-of-channel (BMOC) bugs, deadlocks caused by misuse of mutexes (\eg, lock-with-unlock, double lock), races on struct fields, and races due to errors when using the testing package. The technical details of GCatch are presented in Section 3 of our ASPLOS paper [1]. 

## Installation and Demonstration

GCatch leverages Z3 for constraint solving. If you have already installed Z3, you can use the `install.sh` script is to install GCatch. You can also install Z3 together with GCatch using the `installZ3.sh` script. 

After the installation, you can run the `run.sh` script to execute GCatch on a buggy version of gRPC. 

## Directories

1. Directory `analysis` contains static analysis routines shared by multiple checkers (e.g., computing post-dominator). 

2. Directory `checkers` contains code for implementing the checking functionalities.

3. Directory `cmd` is the entry point of GCatch.

4. Directory `config` contains configuration files of GCatch. 

5. Directory `instinfo` contains code for analyzing Go SSA instructions. 

6. Directory `output` contains code for printing the detection results. 

7. Directory `path` contains code for computation conducted on CFG.

8. Directory `ssabuild` contains code for transforming input programs into SSA. 

9. Directory `syncgraph` contains code for interaction with call graph, alias analysis, and Z3. 

10. Directory `testdata` contains a buggy version of gRPC. 

11. Directory `tests` contains toy programs to test traditional checkers. 

12. Directory `tools` contains copies of external packages. 

13. Directory `util` contains utility code shared by different components of GCatch. 

Please refer to the demonstration scripts to figure out how to use GCatch’s code. 


[1] Ziheng Liu, Shuofei Zhu, Boqin Qin, Hao Chen and Linhai Song. “Automatically Detecting and Fixing Concurrency Bugs in Go Software Systems.” In ASPLOS’2020. 
