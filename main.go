package main

// Linter to check for returned interfaces.

// Go lint as reference:
// https://github.com/golang/lint/blob/8f45f776/lint.go

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path"
)

func fmtError(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}

func main() {
	if len(os.Args) < 2 {
		fmtError("filename to required\n")
		os.Exit(1)
	}

	filepath := os.Args[1]

	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmtError("error: %s\n", err)
		os.Exit(2)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path.Base(os.Args[1]), dat, parser.AllErrors)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(3)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("", fset, []*ast.File{f}, nil)
	scope := pkg.Scope()
	//spew.Dump(f)
	ast.Inspect(f, func(n ast.Node) bool {
		ret, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}
		//spew.Dump(ret.Type.Results)
		if ret.Type.Results == nil {
			return false
		}
		for _, field := range ret.Type.Results.List {
			//spew.Dump(field)
			i, ok := field.Type.(*ast.Ident)
			if !ok {
				continue
			}

			ft := scope.Lookup(i.Name)
			//spew.Dump(ft.Type())
			if ft == nil {
				continue
			}

			t := ft.Type()
			if t == nil {
				fmtError("%s:%d type is nil? FIXME\n", os.Args[1], fset.Position(ret.Pos()).Line)
				os.Exit(4)
			}
			u := t.Underlying()
			if !types.IsInterface(u) {
				continue
			}

			//spew.Dump(u)
			if _, ok := u.(*types.Interface); ok {
				fmt.Printf("%s:%d returned interface: %s\n", os.Args[1], fset.Position(ret.Pos()).Line, ft.Name())
			}
		}
		return true
	})
}
