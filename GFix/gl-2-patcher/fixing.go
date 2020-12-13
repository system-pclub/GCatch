package main

//TODO: try to read info from parameters

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type BugItem struct {
	bugNo   int
	bugType string
	src     string
	lineno  int
}

//https://github.com/zupzup/ast-manipulation-example/blob/master/main.go

func getStmtToInsert(lineNoToStmt int, fset *token.FileSet, f *ast.File) ast.Stmt {
	var ret ast.Stmt = nil
	ast.Inspect(f, func(node ast.Node) bool {
		switch x := node.(type) {
		case ast.Stmt:
			if fset.Position(x.Pos()).Line == lineNoToStmt {
				ret = x
				return false
			}
		}
		return true
	})
	if ret == nil {
		panic("didn't find a statement at line number " + strconv.Itoa(lineNoToStmt))
	}
	return ret
}

func makeDeferFunc(stmt ast.Stmt) *ast.DeferStmt {
	list := []ast.Stmt{
		stmt,
	}
	ret := &ast.DeferStmt{
		Defer: 0,
		Call: &ast.CallExpr{
			Fun: &ast.FuncLit{
				Type: &ast.FuncType{
					Func: 0,
					Params: &ast.FieldList{
						Opening: 0,
						List:    nil,
						Closing: 0,
					},
					Results: nil,
				},
				Body: &ast.BlockStmt{
					Lbrace: 0,
					List:   list,
					Rbrace: 0,
				},
			},
			Lparen:   0,
			Args:     nil,
			Ellipsis: 0,
			Rparen:   0,
		},
	}
	return ret
}

func insertBeforeLineNo(lineno int, stmt ast.Stmt, fset *token.FileSet, f *ast.File) {
	visited := false
	ast.Inspect(f, func(node ast.Node) bool {
		var body *ast.BlockStmt = nil
		switch x := node.(type) {
		case *ast.FuncDecl: //the explicit declaration
			body = x.Body
		case *ast.FuncLit:
			body = x.Body
		}
		if body == nil {
			return true
		}
		if fset.Position(body.Lbrace).Line <= lineno && lineno <= fset.Position(body.Rbrace).Line {
			index := 0
			for i, stmt := range body.List {
				if fset.Position(stmt.Pos()).Line == lineno {
					index = i
					visited = true
					break
				}
			}
			if visited {
				newList := make([]ast.Stmt, len(body.List)+1)
				for i := 0; i < index; i++ {
					newList[i] = body.List[i]
				}
				newList[index] = stmt
				for i := index; i < len(body.List); i++ {
					newList[i+1] = body.List[i]
				}
				body.List = newList
			}
			return false
		}
		return true
	})
	if !visited {
		panic("didn't find a statement at line number " + strconv.Itoa(lineno))
	}
}

func patchOnFile(filename string, code string) {
	f, err := os.Create(filename) //os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("1")
		panic(err)
	}

	_, err = f.WriteString(code)
	if err != nil {
		fmt.Println("2")
		f.Close()
		panic(err)
	}
	f.Close()
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	filename := os.Args[1]                             //filename
	lineNoToInsertDefer, _ := strconv.Atoi(os.Args[2]) //the line number to insert the defer operation
	//We insert the code before the line number. i.e., if lineNoToInsertDefer = 123, after insertion,
	//at line 123 is a defer operation.
	var linesToDelete []int
	for _, strLine := range os.Args[3:] { //line numbers to remove
		lineno, _ := strconv.Atoi(strLine)
		linesToDelete = append(linesToDelete, lineno)
	}
	dat, _ := ioutil.ReadFile(filename)
	// src is the input for which we want to print the AST.
	src := string(dat)
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	patch(lineNoToInsertDefer, linesToDelete, fset, f)

	var retbuf strings.Builder
	err = printer.Fprint(&retbuf, fset, f)
	if err != nil {
		//TODO
	}
	patchedCode := retbuf.String()

	//fmt.Println(patchedCode)
	patchOnFile(filename, patchedCode)
	// Print the AST.
	//ast.Print(fset, f)
}

func patch(deferLineNo int, linesToDelete []int, fset *token.FileSet, f *ast.File) {
	stmt := getStmtToInsert(linesToDelete[0], fset, f)
	deferOp := makeDeferFunc(stmt)
	for _, x := range linesToDelete {
		removeAtLineNo(x, fset, f)
	}
	insertBeforeLineNo(deferLineNo, deferOp, fset, f)
}

func removeAtLineNo(lineno int, fset *token.FileSet, f *ast.File) {
	visited := false
	ast.Inspect(f, func(node ast.Node) bool {
		var body *ast.BlockStmt = nil
		switch x := node.(type) {
		case *ast.FuncDecl: //the explicit declaration
			body = x.Body
		case *ast.FuncLit:
			body = x.Body
		}
		if body == nil {
			return true
		}
		if fset.Position(body.Lbrace).Line <= lineno && lineno <= fset.Position(body.Rbrace).Line {
			for i, stmt := range body.List {
				if fset.Position(stmt.Pos()).Line == lineno {
					body.List[i] = &ast.EmptyStmt{
						Semicolon: 0,
						Implicit:  true,
					}
					visited = true
					return false
				}
			}
		}
		return true
	})
	if !visited {
		panic("didn't find a statement at line number " + strconv.Itoa(lineno))
	}
}
