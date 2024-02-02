package visitor

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
	"weicai.zhao.io/script/internal"
	"weicai.zhao.io/tools"
)

func ColumnToField(columns []internal.TableColumn) []*ast.Field {
	var fields = make([]*ast.Field, 0)
	for i := 0; i < len(columns); i++ {
		field := &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent(columnNameToField(columns[i].COLUMN_NAME)),
			},
			Type: tableColumnToAstType(columns[i].DATA_TYPE),
			Tag:  &ast.BasicLit{Value: tableColumnTagValue(columns[i])},
		}
		fields = append(fields, field)
	}

	return fields
}

func columnTypeToToken(typ string) token.Token {
	var t = token.IDENT
	switch typ {
	case "uint8", "int64":
		t = token.INT
	case "float64", "float":
		t = token.FLOAT
	case "string":
		t = token.STRING
	}
	return t
}

func tableColumnTagValue(item internal.TableColumn) string {
	gormTag := []string{
		fmt.Sprintf("column:%s", item.COLUMN_NAME),
		fmt.Sprintf("type:%s", item.COLUMN_TYPE),
	}
	if item.IS_NULLABLE == "YES" {
		gormTag = append(gormTag, "NULL")
	} else {
		gormTag = append(gormTag, "NOT NULL")
	}
	switch item.COLUMN_KEY {
	case "PRI":
		gormTag = append(gormTag, "primaryKey")
	case "UNI":
		gormTag = append(gormTag, "uniqueKey")
	case "MUL":
		gormTag = append(gormTag, "index")
	}
	gormTagValue := strings.Join(gormTag, ";")

	var jsonTagValue = item.COLUMN_NAME
	if item.COLUMN_NAME == "id" {
		jsonTagValue = "-"
	}

	return fmt.Sprintf("`gorm:\"%s\" comment:\"%s\" json:\"%s\"`", gormTagValue, item.COLUMN_COMMENT, jsonTagValue)
}

func tableColumnToAstType(name string) ast.Expr {
	var typ = ""
	switch name {
	case "tinyint":
		typ = "uint8"
	case "int", "bigint":
		typ = "int64"
	case "double", "float":
		typ = "float64"
	case "time":
		typ = "string"
	case "char":
		typ = "string"
	default:
		typ = "string"
	}
	return ast.NewIdent(typ)
}

func columnNameToField(s string) string {
	if s == "id" || s == "uuid" {
		return strings.ToUpper(s)
	}

	return tools.UnderlineToUpperCamelCase(s)
}
