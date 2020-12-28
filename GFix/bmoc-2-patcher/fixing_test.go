package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
	"testing"
)

var src = `package main

func gl1_before() {
	done := make(chan struct{})
	go func() {
		<-done
	}()
	done <- struct{}{}
	return
}
`
var lineno = 8

/*var expected = `
`*/

func TestRemove(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	removeAtLineNo(lineno, fset, f)
	var retbuf strings.Builder
	err = printer.Fprint(&retbuf, fset, f)
	if err != nil {
		//TODO
		println(err)
	}
	patchedCode := retbuf.String()
	print(patchedCode)
	ast.Print(fset, f)
}

func TestInsert(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	stmt := getStmtToInsert(lineno, fset, f)
	deferOp := makeDeferFunc(stmt)
	insertBeforeLineNo(4, deferOp, fset, f)
	var retbuf strings.Builder
	err = printer.Fprint(&retbuf, fset, f)
	if err != nil {
		//TODO
		println(err)
	}
	patchedCode := retbuf.String()
	print(patchedCode)
	ast.Print(fset, f)
}

func TestInsertAndRemove(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	patch(4, []int{lineno}, fset, f)
	var retbuf strings.Builder
	err = printer.Fprint(&retbuf, fset, f)
	if err != nil {
		//TODO
		println(err)
	}
	patchedCode := retbuf.String()
	print(patchedCode)
	ast.Print(fset, f)
}
