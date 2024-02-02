package tools

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func ParseAstFile(filename string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}

	return fset, f, nil
}
