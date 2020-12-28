package main

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

func getLineNoOfMakeChan(lineNoToStmt int, fset *token.FileSet, f *ast.File) int {
	ret := 0
	isInVarDecl := false
	ast.Inspect(f, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.DeclStmt:
			decl, ok := x.Decl.(*ast.GenDecl)
			if ok {
				if fset.Position(decl.Lparen).Line <= lineNoToStmt && lineNoToStmt <= fset.Position(decl.Rparen).Line {
					isInVarDecl = true
					ret = fset.Position(decl.Rparen).Line + 1
					return false
				}
			}
		case ast.Stmt:
			if fset.Position(x.Pos()).Line == lineNoToStmt {
				if !isInVarDecl {
					ret = lineNoToStmt
				}
				return false
			}
		}
		return true
	})
	if ret == 0 {
		panic("didn't find a statement at line number " + strconv.Itoa(lineNoToStmt))
	}
	return ret
}

func makeMakeChanAndDefer(name string) (*ast.AssignStmt, *ast.DeferStmt) {
	assign := &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.Ident{
				NamePos: 0,
				Name:    name,
				Obj:     ast.NewObj(ast.Var, name),
			},
		},
		TokPos: 0,
		Tok:    token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun:    ast.NewIdent("make"),
				Lparen: 0,
				Args: []ast.Expr{
					&ast.ChanType{
						Begin: 0,
						Arrow: token.NoPos,
						Dir:   3,
						Value: &ast.StructType{
							Struct: 0,
							Fields: &ast.FieldList{
								Opening: 0,
								List:    nil,
								Closing: 0,
							},
							Incomplete: false,
						},
					},
				},
				Ellipsis: token.NoPos,
				Rparen:   0,
			},
		},
	}
	deferstmt := &ast.DeferStmt{
		Defer: 0,
		Call: &ast.CallExpr{
			Fun:    ast.NewIdent("close"),
			Lparen: 0,
			Args: []ast.Expr{
				ast.NewIdent(name),
			},
			Ellipsis: token.NoPos,
			Rparen:   0,
		},
	}
	return assign, deferstmt
}

func makeSelect(doneChannelName string, buggyStmt ast.Stmt) *ast.SelectStmt {
	selectStmt := &ast.SelectStmt{
		Select: 0,
		Body: &ast.BlockStmt{
			Lbrace: 0,
			List: []ast.Stmt{
				&ast.CommClause{
					Case: 0,
					Comm: &ast.ExprStmt{X: &ast.UnaryExpr{
						OpPos: 0,
						Op:    token.ARROW,
						X:     ast.NewIdent(doneChannelName),
					}},
					Colon: 0,
					Body: []ast.Stmt{
						&ast.ReturnStmt{
							Return:  0,
							Results: nil,
						},
					},
				},
				&ast.CommClause{
					Case:  0,
					Comm:  buggyStmt,
					Colon: 0,
					Body:  nil,
				},
			},
			Rbrace: 0,
		},
	}
	return selectStmt
}

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

func insertBeforeLineNo(lineno int, stmts []ast.Stmt, fset *token.FileSet, f *ast.File) {
	minLineNo := 99999999
	var outerBody *ast.BlockStmt
	index := 0
	ast.Inspect(f, func(node ast.Node) bool {
		var body *ast.BlockStmt = nil
		switch x := node.(type) {
		case *ast.BlockStmt:
			body = x
			/*		case *ast.FuncDecl: //the explicit declaration
						body = x.Body
					case *ast.FuncLit:
						body = x.Body*/
		}
		if body == nil {
			return true
		}
		if fset.Position(body.Lbrace).Line <= lineno && lineno <= fset.Position(body.Rbrace).Line {
			for i, stmt := range body.List {
				//println(fset.Position(stmt.Pos()).Line)
				currentLineNo := fset.Position(stmt.Pos()).Line
				if currentLineNo >= lineno {
					if currentLineNo < minLineNo {
						minLineNo = currentLineNo
						index = i
						outerBody = body
					}
				}
			}
			return true
		}
		return true
	})
	if minLineNo != 99999999 {
		l := len(stmts)
		newList := make([]ast.Stmt, len(outerBody.List)+l)
		for i := 0; i < index; i++ {
			newList[i] = outerBody.List[i]
		}
		//newList[index] = stmts[]
		for i := 0; i < l; i++ {
			newList[index+i] = stmts[i]
		}
		for i := index; i < len(outerBody.List); i++ {
			newList[i+l] = outerBody.List[i]
		}
		outerBody.List = newList
	} else {
		panic("didn't find a statement at line number " + strconv.Itoa(lineno))
	}
}

func replaceAtLineNo(lineno int, stmtToReplace ast.Stmt, fset *token.FileSet, f *ast.File) {
	visited := false
	ast.Inspect(f, func(node ast.Node) bool {
		var body *ast.BlockStmt = nil
		switch x := node.(type) {
		case *ast.BlockStmt:
			body = x
			/*case *ast.FuncDecl: //the explicit declaration
				body = x.Body
			case *ast.FuncLit:
				body = x.Body*/
		}
		if body == nil {
			return true
		}
		if fset.Position(body.Lbrace).Line <= lineno && lineno <= fset.Position(body.Rbrace).Line {
			for i, stmt := range body.List {
				if fset.Position(stmt.Pos()).Line == lineno {
					body.List[i] = stmtToReplace
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

func patch(makeDeferLineNo int, linesToDelete []int, fset *token.FileSet, f *ast.File) {
	doneChannelName := "__auto_patched_stop"
	makechanOp, deferOp := makeMakeChanAndDefer(doneChannelName)
	for _, x := range linesToDelete {
		stmt := getStmtToInsert(x, fset, f)
		selectStmt := makeSelect(doneChannelName, stmt)
		replaceAtLineNo(x, selectStmt, fset, f)
	}
	insertBeforeLineNo(makeDeferLineNo, []ast.Stmt{makechanOp, deferOp}, fset, f)
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

//params: filename, lineno to insert make chan and defer, lineno to rewrite as select
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
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	lineNoToInsertDefer = getLineNoOfMakeChan(lineNoToInsertDefer, fset, f)
	patch(lineNoToInsertDefer, linesToDelete, fset, f)

	var retbuf strings.Builder
	err = printer.Fprint(&retbuf, fset, f)
	if err != nil {
		//TODO
	}
	patchedCode := retbuf.String()

	fmt.Println(patchedCode)
	//patchOnFile(filename, patchedCode)
	// Print the AST.
	//ast.Print(fset, f)
}
