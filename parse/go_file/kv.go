package go_file

import (
	"fmt"
	"strings"
)

type Param struct {
	Names []string
	Type  string
}

func (p Param) ToString() string {
	if p.Names == nil || len(p.Names) == 0 {
		return p.Type
	}
	return fmt.Sprintf("%s %s", strings.Join(p.Names, ", "), p.Type)
}

type Params []Param

func (p Params) ToString() string {
	var s = make([]string, 0)
	for i := 0; i < len(p); i++ {
		s = append(s, p[i].ToString())
	}
	return fmt.Sprintf("(%s)", strings.Join(s, ", "))
}

type Return Param

func (p Return) ToString() string {
	if p.Names == nil || len(p.Names) == 0 {
		return p.Type
	}
	return fmt.Sprintf("%s %s", strings.Join(p.Names, ", "), p.Type)
}

type Returns []Return

func (p Returns) ToString() string {
	if p == nil || len(p) == 0 {
		return ""
	}
	var s = make([]string, 0)
	for i := 0; i < len(p); i++ {
		s = append(s, p[i].ToString())
	}
	if len(s) == 1 {
		return s[0]
	}
	return fmt.Sprintf("(%s)", strings.Join(s, ", "))
}

type Owner struct {
	Usage string
	Name  string
}

func (o Owner) ToString() string {
	if o.Usage == "" {
		return fmt.Sprintf("(%s)", o.Name)
	}
	return fmt.Sprintf("(%s %s)", o.Usage, o.Name)
}

type Const struct {
	Name  string
	Value string
}

func (c Const) ToString() string {
	return fmt.Sprintf("const %s = %s", c.Name, c.Value)
}

type Consts []Const

func (c Consts) ToString() string {
	if c == nil || len(c) == 0 {
		return ""
	}
	if len(c) == 1 {
		return c[0].ToString()
	}
	var res = make([]string, 0)
	for i := 0; i < len(c); i++ {
		res = append(res, strings.ReplaceAll(c[i].ToString(), "const ", "    "))
	}
	return fmt.Sprintf("const (\n%s\n)", strings.Join(res, "\n"))
}

type Var struct {
	Name  string
	Value string
}

func (v Var) ToString() string {
	return fmt.Sprintf("var %s = %s", v.Name, v.Value)
}

type Vars []Var

func (v Vars) ToString() string {
	if v == nil || len(v) == 0 {
		return ""
	}
	if len(v) == 1 {
		return v[0].ToString()
	}
	var res = make([]string, 0)
	for i := 0; i < len(v); i++ {
		res = append(res, strings.ReplaceAll(v[i].ToString(), "var ", "    "))
	}
	return fmt.Sprintf("var (\n%s\n)", strings.Join(res, "\n"))
}

type Tag interface {
	ToString() string
}

type JsonTag struct {
	Key   string
	Value []string
}

func (t JsonTag) ToString() string {
	return fmt.Sprintf(`json:"%s"`, strings.Join(t.Value, ","))
}

type GormTag struct {
	Key   string
	Value []string
}

func (t GormTag) ToString() string {
	return fmt.Sprintf(`gorm:"%s"`, strings.Join(t.Value, ";"))
}

type Tags []Tag

func (t Tags) ToString() string {
	if t == nil || len(t) == 0 {
		return ""
	}
	var res = make([]string, 0)
	for i := 0; i < len(t); i++ {
		res = append(res, t[i].ToString())
	}
	return fmt.Sprintf("`%s`", strings.Join(res, " "))
}

type Import struct {
	Usage string
	Value string
}

func (i Import) ToString() string {
	if i.Usage != "" {
		return fmt.Sprintf(`%s "%s"`, i.Usage, i.Value)
	}
	return fmt.Sprintf(`"%s"`, i.Value)
}

type Imports []Import

func (i Imports) GetImports() Imports {
	var res = Imports{}
	if i == nil || len(i) == 0 {
		return res
	}
	var exist = make(map[string]bool)
	for j := 0; j < len(i); j++ {
		if i[j].Value == "" {
			continue
		}
		if ok, _ := exist[i[j].Value]; ok {
			continue
		}
		res = append(res, i[j])
		exist[i[j].Value] = true
	}
	return res
}

func (i Imports) ToString() string {
	im := i.GetImports()
	if im == nil || len(im) == 0 {
		return ""
	}
	var res = make([]string, 0)
	for j := 0; j < len(im); j++ {
		res = append(res, "    "+im[j].ToString())
	}
	return fmt.Sprintf("import (\n%s\n)", strings.Join(res, "\n"))
}
