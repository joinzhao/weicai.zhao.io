package go_file

import (
	"fmt"
	"strings"
)

type Function struct {
	Imports      Imports
	FunctionName string
	Params       Params
	Owner        Owner
	Returns      Returns
	Value        string
}

func (f Function) GetImports() Imports {
	return f.Imports
}

func (f Function) ToString() string {
	var tmp = "func"
	if f.Owner.Name != "" {
		tmp += " " + f.Owner.ToString()
	}

	tmp += fmt.Sprintf(" %s", f.FunctionName)

	if f.Params != nil {
		tmp += " " + f.Params.ToString()
	} else {
		tmp += " ()"
	}

	if f.Returns != nil {
		tmp += " " + f.Returns.ToString()
	}

	if f.Value == "" {
		tmp += " {}"
	} else {
		tmp += fmt.Sprintf(" {\n%s\n}", f.Value)
	}

	return tmp
}

type Functions []Function

func (f Functions) GetImports() Imports {
	var res = Imports{}
	if f == nil || len(f) == 0 {
		return res
	}
	for i := 0; i < len(f); i++ {
		im := f[i].GetImports()
		if im != nil && len(im) > 0 {
			res = append(res, im...)
		}
	}
	return res.GetImports()
}

func (f Functions) ToString() string {
	var s = make([]string, 0)
	for i := 0; i < len(f); i++ {
		s = append(s, f[i].ToString())
	}
	return fmt.Sprintf("%s", strings.Join(s, "\n"))
}
