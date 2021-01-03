Automatically Detecting and Fixing Concurrency Bugs in Go Software Systems

## Descriptions

This is the code repository of our ASPLOS paper [1]. GCatch is a suite of static detectors that can analyze large, real Go software. GFix is an automated fixing tool that can synthesize patches for blocking misuse-of-channel (BMOC) bugs detected by GCatch. We evaluated GCatch and GFix in 21 open-source Go projects (e.g., Docker, Kubernetes, gRPC). In total, GCatch detects XXX BMOC bugs and XXX traditional concurrency bugs and GFix successfully generates patches for XXX BMOC bugs. The detailed experimental data can be found [here](https://docs.google.com/spreadsheets/d/1mDxB6IRxrTodF9CrmpUu72E6673y5s9BkjKuTjtx1qc/edit#gid=1005528908). 


[1] Ziheng Liu, Shuofei Zhu, Boqin Qin, Hao Chen, and Linhai Song. “Automatically Detecting and Fixing Concurrency Bugs in Go Software Systems.” In ASPLOS’2020. 

