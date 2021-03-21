## Descriptions of tools

GCatch depends on several libraries. In order to change these libraries and
avoid the update of libraries influences the functionality of GCatch, 
in `tools` we forked these libraries.

**Libraries**
1. golang.org/x/tools/go
   
   _Forked around: 07/2019_
   
   This library contains useful functions for static analysis. 
   GCatch mainly changed the following parts of this library:
   building of AST and SSA, pointer analysis, call-graph generation.

2. github.com/aclements/go-z3
   
   _Forked at commit: 18129d7fc68746d95902a76704a19f68490c7ebf_
   
   This library provides useful Golang's wrappers of z3 libraries in C. 
   GCatch doesn't change this library.
   
3. z3 

   _Forked around: 01/2021_

    BMOC checker (and go-z3) depends on Z3 to solve the constraints 
    and detect BMOC bugs.
    GCatch doesn't change this library.