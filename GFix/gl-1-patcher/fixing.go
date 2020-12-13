package main
//TODO: try to read info from parameters

import (
	"encoding/csv"
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
	bugNo int
	bugType string
	src string
	lineno int
}

//https://github.com/zupzup/ast-manipulation-example/blob/master/main.go
func fixGoroutineLeakOnChannelType1(lineno int, fset *token.FileSet, f *ast.File) {
	var argSite *ast.CallExpr = nil
	ast.Inspect(f, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.CallExpr:
			var fun = x.Fun
			switch y := fun.(type) {
			case *ast.Ident:
				if y.Name == "make" && fset.Position(y.Pos()).Line == lineno{//
					fmt.Println("found 'make' in the line")
					if len(x.Args) == 1 {
						argSite = x
					}
					//argSite = x.Args
					//x.Args = append(x.Args, &ast.BasicLit{Kind: token.INT, Value: "1"})
					//would be better to insert after visit
				}
			}
		}
		return true
	})
	if argSite != nil {
		argSite.Args = append(argSite.Args, &ast.BasicLit{Kind: token.INT, Value: "1"})
	}

	//ast.Print(fset, f)
}

func readBugList(listPath string) map[int]*BugItem {
	fmt.Println(listPath)
	fd, _ := os.Open(listPath)//"/home/suz305/go-bug-study/projects/src/gl_patcher/bug_list.txt")
	defer fd.Close()
	rd := csv.NewReader(fd)
	lines, err := rd.ReadAll()
	if err != nil {
		fmt.Println("error reading the file")
		fmt.Println(listPath)
		return nil
	}
	print(len(lines))
	bugs := make(map[int]*BugItem, len(lines))
	for _, line := range lines[1:] {
		bugno, _ := strconv.Atoi(line[0])
		bugtype := line[1]
		src := line[2]
		lineno, _ := strconv.Atoi(line[3])
		bugs[bugno] = &BugItem{
			bugNo:   bugno,
			bugType: bugtype,
			src:     src,
			lineno:  lineno,
		}
	}
	return bugs
}


func patch(filename string, code string) {
	f, err := os.Create(filename)//os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)

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
	filename:= os.Args[1]//filename
	//lineno, _ := strconv.Atoi(os.Args[2])//lineno
	dat, _ := ioutil.ReadFile(filename)
	// src is the input for which we want to print the AST.
	src := string(dat)
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	for _, strLine := range os.Args[2:] {
		lineno, _ := strconv.Atoi(strLine)
		fixGoroutineLeakOnChannelType1(lineno, fset, f)
	}

	var retbuf strings.Builder
	err = printer.Fprint(&retbuf, fset, f)
	if err != nil {
		//TODO
	}
	patchedCode := retbuf.String()

	//fmt.Println(patchedCode)
	patch(filename, patchedCode)
	// Print the AST.
	//ast.Print(fset, f)
}
