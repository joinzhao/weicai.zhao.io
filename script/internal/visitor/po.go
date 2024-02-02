package visitor

import (
	"fmt"
	"go/ast"
	"go/token"
)

type poVisitor struct {
	fields     []*ast.Field
	structName string
	tableName  string
}

func NewPoVisitor(fields []*ast.Field, structName, tableName string) ast.Visitor {
	return &poVisitor{
		fields:     fields,
		structName: structName,
		tableName:  tableName,
	}
}

func (v *poVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.GenDecl:
		v.gen(node.(*ast.GenDecl))
	case *ast.FuncDecl:
		v.funcDecl(node.(*ast.FuncDecl))
	case *ast.StructType:
		v.structType(node.(*ast.StructType))
	}

	return v
}

func (v *poVisitor) gen(node *ast.GenDecl) {
	switch node.Tok {
	case token.CONST:
		v.genConst(node.Specs[0].(*ast.ValueSpec))
	case token.TYPE:
		v.genType(node.Specs[0].(*ast.TypeSpec))
	}
}
func (v *poVisitor) genConst(node *ast.ValueSpec) {
	node.Names = []*ast.Ident{
		ast.NewIdent(v.structName + "TableName"),
	}
	node.Values = []ast.Expr{
		&ast.BasicLit{
			Value: fmt.Sprintf(`"%s"`, v.tableName),
			Kind:  token.STRING,
		},
	}
}

func (v *poVisitor) genType(node *ast.TypeSpec) {
	node.Name = ast.NewIdent(v.structName)
}

func (v *poVisitor) structType(node *ast.StructType) {
	node.Fields = &ast.FieldList{
		List: v.fields,
	}
}

func (v *poVisitor) funcDecl(node *ast.FuncDecl) {
	node.Recv = &ast.FieldList{
		Opening: 0,
		List: []*ast.Field{
			{
				Type: ast.NewIdent(v.structName),
			},
		},
		Closing: 0,
	}

	node.Body.List[0].(*ast.ReturnStmt).Results = []ast.Expr{
		&ast.Ident{
			Name: v.structName + "TableName",
		},
	}
}
