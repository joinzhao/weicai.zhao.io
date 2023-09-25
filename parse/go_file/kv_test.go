package go_file

import "testing"

func TestParam_ToString(t *testing.T) {
	var p = Param{
		Names: []string{},
		Type:  "string",
	}
	t.Log(p.ToString())

	p = Param{
		Names: []string{"p1", "p2"},
		Type:  "string",
	}
	t.Log(p.ToString())
}

func TestParams_ToString(t *testing.T) {
	var p = Params{
		{
			Names: []string{},
			Type:  "string",
		},
	}
	t.Log(p.ToString())

	p = append(p, Param{
		Names: []string{"p1", "p2"},
		Type:  "string",
	},
	)
	t.Log(p.ToString())
}

func TestOwner_ToString(t *testing.T) {
	var o = Owner{
		Usage: "f",
		Name:  "file",
	}
	t.Log(o.ToString())

	o = Owner{
		Usage: "",
		Name:  "file",
	}
	t.Log(o.ToString())

	o = Owner{
		Usage: "f",
		Name:  "*file",
	}
	t.Log(o.ToString())

	o = Owner{
		Usage: "",
		Name:  "*file",
	}
	t.Log(o.ToString())
}

func TestConst_ToString(t *testing.T) {
	var c = Const{
		Name:  "key",
		Value: `"val"`,
	}

	t.Log(c.ToString())
}

func TestConsts_ToString(t *testing.T) {
	var c = Consts{}

	t.Log(c.ToString())

	c = append(c, Const{
		Name:  "key",
		Value: `"val"`,
	})
	t.Log(c.ToString())

	c = append(c, Const{
		Name:  "key1",
		Value: `"val1"`,
	})
	t.Log(c.ToString())
}

func TestVar_ToString(t *testing.T) {
	var c = Var{
		Name:  "key",
		Value: `"val"`,
	}

	t.Log(c.ToString())
}

func TestVars_ToString(t *testing.T) {
	var c = Vars{}

	t.Log(c.ToString())

	c = append(c, Var{
		Name:  "key",
		Value: `"val"`,
	})
	t.Log(c.ToString())

	c = append(c, Var{
		Name:  "key1",
		Value: `"val1"`,
	})
	t.Log(c.ToString())
}

func TestJsonTag_ToString(t *testing.T) {
	var j = JsonTag{
		Key:   "",
		Value: []string{"t1", "t2"},
	}
	t.Log(j.ToString())
}

func TestGormTag_ToString(t *testing.T) {
	var g = GormTag{
		Key:   "",
		Value: []string{"t1", "t2"},
	}
	t.Log(g.ToString())
}

func TestTags_ToString(t *testing.T) {
	var g = Tags{
		JsonTag{
			Key:   "",
			Value: []string{"t1", "t2"},
		},
		GormTag{
			Key:   "",
			Value: []string{"t1", "t2"},
		},
	}
	t.Log(g.ToString())
}

func TestImport_ToString(t *testing.T) {
	var im = Import{
		Usage: ".",
		Value: "gorm.io/gorm",
	}
	t.Log(im.ToString())

	im.Usage = ""
	t.Log(im.ToString())
}

func TestImports_GetImports(t *testing.T) {
	var ims = Imports{}
	t.Log(ims.GetImports())

	ims = append(ims, Import{
		Usage: ".",
		Value: "gorm.io/gorm",
	})
	t.Log(ims.GetImports())

	ims = append(ims, Import{
		Usage: "",
		Value: "gorm.io/gorm",
	})
	t.Log(ims.GetImports())
}

func TestImports_ToString(t *testing.T) {
	var ims = Imports{}
	t.Log(ims.ToString())

	ims = append(ims, Import{
		Usage: ".",
		Value: "gorm.io/gorm",
	})
	t.Log(ims.ToString())

	ims = append(ims, Import{
		Usage: "",
		Value: "gorm.io/gorm",
	})
	t.Log(ims.ToString())

	ims = append(ims, Import{
		Usage: ".",
		Value: "weicai.zhao.io",
	})
	t.Log(ims.ToString())
}
