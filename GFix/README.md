# GFix: Automatically Fixing Blocking Misuse-of-Channel Bugs 

## Descriptions

GFix aims to automatically fix blocking misuse-of-channel (BMOC) bugs detected by GCatch. It implements three fixing strategies and aims to generate patches with good performance and readability. GFix takes two steps for each input bug. It first applies its dispatcher component to decide whether an input bug can be fixed and if so, which fixing strategy to use. It then invokes one of the three patchers to generate a concrete patch. The technical details of GFix are presented in Section 4 and evaluation results are discussed in Section 5.3 of our ASPLOS paper [1].


## Examples and Demonstrations

1. In the `etcd-example` directory, the `run.sh` script is to download the source code of etcd, change it to a buggy version, and generate three patches using the three fixing strategies for three different bugs detected by GCatch. 

2. In `toy-examples/src` directory, there are three toy programs to demonstrate the three types of bugs that can be fixed by GFix. The `run.sh` script is to apply GFix to the three bugs contained in the three programs. 

## Directories

1. Directory `dispatcher` contains related code for the dispatcher component. 

2. Directory `gl-1-patcher` contains code for generating patches using the first fixing strategy. 

3. Directory `gl-2-patcher` contains code for generating patches using the second fixing strategy. 

4. Directory `gl-3-patcher` contains code for generating patches using the third fixing strategy. 

Please refer to the demonstration scripts to figure out how to use GFix’s code. 


[1] Ziheng Liu, Shuofei Zhu, Boqin Qin, Hao Chen and Linhai Song. “Automatically Detecting and Fixing Concurrency Bugs in Go Software Systems.” In ASPLOS’2020. 

