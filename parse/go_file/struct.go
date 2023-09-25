package go_file

import (
	"fmt"
	"strings"
)

type Struct struct {
	StructName string
	Usage      string
	Comment    string
	Fields     Fields
	Functions  Functions
}

func (s Struct) GetImports() Imports {
	var res = Imports{}
	if s.Fields != nil {
		im := s.Fields.GetImports()
		if im != nil && len(im) > 0 {
			res = append(res, im...)
		}
	}
	if s.Functions != nil {
		im := s.Functions.GetImports()
		if im != nil && len(im) > 0 {
			res = append(res, im...)
		}
	}

	return res.GetImports()
}

func (s Struct) ToString() string {
	var tmp = ""
	var args = make([]any, 0)
	if s.Comment != "" {
		tmp += "// %s %s \n"
		args = append(args, s.StructName, s.Comment)
	}
	tmp += "type %s struct{\n%s\n}\n\n"
	args = append(args, s.StructName, s.Fields.ToString())
	var fs = s.Functions
	if fs != nil {
		for i := 0; i < len(fs); i++ {
			fs[i].Owner.Name = fmt.Sprintf("*%s", s.StructName)
			fs[i].Owner.Usage = s.Usage
		}
		tmp += "%s\n\n"
		args = append(args, fs.ToString())
	}

	return fmt.Sprintf(tmp, args...)
}

type Structs []Struct

func (s Structs) GetImports() Imports {
	var res = Imports{}
	if s == nil || len(s) == 0 {
		return res
	}
	for i := 0; i < len(s); i++ {
		im := s[i].GetImports()
		if im != nil || len(im) > 0 {
			res = append(res, im...)
		}
	}
	return res.GetImports()
}

func (s Structs) ToString() string {
	if s == nil || len(s) == 0 {
		return ""
	}
	var res = make([]string, 0)
	for i := 0; i < len(s); i++ {
		res = append(res, s[i].ToString())
	}
	return strings.Join(res, "\n\n")
}
