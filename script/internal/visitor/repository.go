package visitor

import (
	"fmt"
	"go/ast"
	"go/token"
	"weicai.zhao.io/script/internal"
	"weicai.zhao.io/tools"
)

func NewRepoVisitor(repoName, newPoPath, oldPoPath string, column internal.TableColumn) ast.Visitor {
	return &repositoryVisitor{
		repoName:  repoName,
		columns:   column,
		newPoPath: newPoPath,
		oldPoPath: oldPoPath,
	}
}

func NewDependencyVisitor(repoName, newPoPath, oldPoPath string, column internal.TableColumn) ast.Visitor {
	return &repositoryVisitor{
		repoName:  repoName,
		columns:   column,
		newPoPath: newPoPath,
		oldPoPath: oldPoPath,
	}
}

type repositoryVisitor struct {
	repoName  string
	columns   internal.TableColumn
	newPoPath string
	oldPoPath string
}

func (v *repositoryVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.KeyValueExpr:
		if node.(*ast.KeyValueExpr).Key.(*ast.Ident).Name == "Column" {
			node.(*ast.KeyValueExpr).Value = &ast.BasicLit{
				ValuePos: 0,
				Kind:     columnTypeToToken(v.columns.DATA_TYPE),
				Value:    fmt.Sprintf(`"%s"`, v.columns.COLUMN_NAME),
			}
		}
	case *ast.Field:
		var field = node.(*ast.Field)
		if field.Names != nil && len(field.Names) > 0 {
			column := tools.UnderlineToLowerCamelCase(v.columns.COLUMN_NAME)
			if field.Names[0].Name == "id" || field.Names[0].Name == column {
				field.Type = tableColumnToAstType(v.columns.DATA_TYPE)
			}
		}
	case *ast.SelectorExpr:
		if node.(*ast.SelectorExpr).Sel.Name == "Model" {
			node.(*ast.SelectorExpr).Sel.Name = v.repoName
			return nil
		}
	case *ast.Ident:
		if node.(*ast.Ident).Name == "id" {
			node.(*ast.Ident).Name = tools.UnderlineToLowerCamelCase(v.columns.COLUMN_NAME)
		}
		if node.(*ast.Ident).Name == "ModelRepo" {
			node.(*ast.Ident).Name = fmt.Sprintf("%sRepo", v.repoName)
		}
		if node.(*ast.Ident).Name == "NewModelRepo" {
			node.(*ast.Ident).Name = fmt.Sprintf("New%sRepo", v.repoName)
		}
		if node.(*ast.Ident).Name == "Model" {
			node.(*ast.Ident).Name = fmt.Sprintf("%sModel", v.repoName)
		}
	case *ast.ImportSpec:
		if v.newPoPath != "" {
			var val = node.(*ast.ImportSpec).Path.Value
			if val == fmt.Sprintf(`"%s"`, v.oldPoPath) {
				node.(*ast.ImportSpec).Path = &ast.BasicLit{
					ValuePos: 0,
					Kind:     token.STRING,
					Value:    fmt.Sprintf(`"%s"`, v.newPoPath),
				}
			}
		}
	}
	return v
}
