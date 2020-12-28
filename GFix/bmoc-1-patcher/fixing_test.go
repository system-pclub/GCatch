package main

import (
	"go/parser"
	"go/token"
	"testing"
)

var src = `package main

func gl1_before() {
	done := make(chan struct{})
	go func() {
		done <- struct{}{}
	}()
	<-done
	return
}
`
var lineno = 4

var expected = `package main

func gl1_before() {
	done := make(chan struct{}, 1)
	go func() {
		done <- struct{}{}
	}()
	<-done
	return
}
`

func TestFixGoroutineLeakOnChannelType1(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	fixGoroutineLeakOnChannelType1(lineno, fset, f)
	/*if patchedCode != expected  {
		t.Fatal("failed")
	}*/
}
