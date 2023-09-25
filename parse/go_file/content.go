package go_file

import "fmt"

type Content struct {
	PackageName string
	Imports     Imports
	Consts      Consts
	Vars        Vars
	Structs     Structs
	Functions   Functions
}

func (c Content) GetImports() Imports {
	var res = Imports{}
	if c.Imports != nil {
		res = c.Imports
	}
	if c.Structs != nil {
		for i := 0; i < len(c.Structs); i++ {
			im := c.Structs.GetImports()
			if im != nil || len(im) > 0 {
				res = append(res, im...)
			}
		}
	}

	if c.Functions != nil {
		for i := 0; i < len(c.Functions); i++ {
			im := c.Functions.GetImports()
			if im != nil || len(im) > 0 {
				res = append(res, im...)
			}
		}
	}
	return res.GetImports()
}

func (c Content) ToString() string {
	var args = make([]any, 0)
	var tmp = "package %s\n\n"
	args = append(args, c.PackageName)

	var im = c.GetImports()
	if im != nil && len(im) > 0 {
		tmp += "%s\n\n"
		args = append(args, im.ToString())
	}
	if c.Consts != nil && len(c.Consts) > 0 {
		tmp += "%s\n\n"
		args = append(args, c.Consts.ToString())
	}

	if c.Vars != nil && len(c.Vars) > 0 {
		tmp += "%s\n\n"
		args = append(args, c.Vars.ToString())
	}

	if c.Structs != nil && len(c.Structs) > 0 {
		tmp += "%s\n\n"
		args = append(args, c.Structs.ToString())
	}

	if c.Functions != nil && len(c.Functions) > 0 {
		tmp += "%s\n\n"
		args = append(args, c.Functions.ToString())
	}

	return fmt.Sprintf(tmp, args...)
}
