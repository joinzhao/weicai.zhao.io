package go_file

import (
	"fmt"
	"strings"
)

type Field struct {
	Import    Import
	FieldName string
	Type      string
	Tags      Tags
	Comment   string
}

func (f Field) GetImport() Import {
	return f.Import
}

func (f Field) ToString() string {
	if f.Comment != "" {
		return fmt.Sprintf("%s %s %s // %s", f.FieldName, f.Type, f.Tags.ToString(), f.Comment)
	}
	return fmt.Sprintf("%s %s %s", f.FieldName, f.Type, f.Tags.ToString())
}

type Fields []Field

func (f Fields) ToString() string {
	if f == nil {
		return ""
	}
	var res = make([]string, 0)
	if f == nil || len(f) == 0 {
		return ""
	}
	for i := 0; i < len(f); i++ {
		res = append(res, "    "+f[i].ToString())
	}
	return fmt.Sprintf("%s", strings.Join(res, "\n"))
}

func (f Fields) GetImports() Imports {
	var res = Imports{}
	if f == nil || len(f) == 0 {
		return res
	}

	for i := 0; i < len(f); i++ {
		res = append(res, f[i].GetImport())
	}
	return res.GetImports()
}
