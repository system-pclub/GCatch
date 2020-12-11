package mypointer

import (
	"github.com/system-pclub/GCatch/GCatch/tools/go/ssa"
	"strings"
)

// There are a few function calls we'd like to handle context sensitively

// 1. context APIs
// context APIs
var ctx_fn_list = []string{"context.WithValue","context.WithCancel","context.WithDeadline","context.WithTimeout"}
// some people still like to import from golang.org/
var ctx_fn_contain_list = []string{"golang.org/x/net/context.WithValue","golang.org/x/net/context.WithCancel","golang.org/x/net/context.WithDeadline","golang.org/x/net/context.WithTimeout"}

// 2. special functions
var special_fn_name_list = []string{"mygo"}

func is_in_sensitive_list(fn *ssa.Function) bool {
	fn_str := fn.String()
	fn_name := fn.Name()
	if is_ctx_fn(fn_str) || is_special_fn_name(fn_name) {
		return true
	}
	return false
}

func is_special_fn_name(name string) bool {
	for _,sp_name := range special_fn_name_list {
		if sp_name == name {
			return true
		}
	}
	return false
}

func is_ctx_fn(str string) bool {
	for _,fn_str := range ctx_fn_list {
		if str == fn_str {
			return true
		}
	}
	for _,fn_str := range ctx_fn_contain_list {
		if strings.Contains(str,fn_str) {
			return true
		}
	}
	return false
}

