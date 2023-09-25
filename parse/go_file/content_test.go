package go_file

import "testing"

func TestContent_GetImports(t *testing.T) {
	var c = Content{
		Imports: Imports{
			{
				Usage: ".",
				Value: "gorm.io/gorm",
			},
			{
				Usage: "",
				Value: "weicai.zhao.io/tools",
			},
			{
				Usage: "",
				Value: "weicai.zhao.io/typex",
			},
		},
		Structs: Structs{
			{
				StructName: "",
				Usage:      "",
				Comment:    "",
				Fields: Fields{
					{
						Import: Import{
							Usage: "",
							Value: "weicai.zhao.io/typex",
						},
					},
					{
						Import: Import{
							Usage: ".",
							Value: "gorm.io/gorm",
						},
					},
					{
						Import: Import{
							Usage: "",
							Value: "weicai.zhao.io/tools",
						},
					},
				},
				Functions: Functions{
					{
						Imports: Imports{
							{
								Usage: "",
								Value: "weicai.zhao.io/typex",
							},
						},
					},
					{
						Imports: Imports{
							{
								Usage: ".",
								Value: "gorm.io/gorm",
							},
						},
					},
					{
						Imports: Imports{
							{
								Usage: "",
								Value: "weicai.zhao.io/tools",
							},
						},
					},
				},
			},
		},
		Functions: Functions{
			{
				Imports: Imports{
					{
						Usage: "",
						Value: "weicai.zhao.io/typex",
					},
				},
			},
			{
				Imports: Imports{
					{
						Usage: ".",
						Value: "gorm.io/gorm",
					},
				},
			},
			{
				Imports: Imports{
					{
						Usage: "",
						Value: "weicai.zhao.io/tools",
					},
				},
			},
		},
	}

	t.Log(c.GetImports())
}

func TestContent_ToString(t *testing.T) {
	var c = Content{
		PackageName: "po",
	}
	t.Log(c.ToString())

	c.Imports = Imports{
		{
			Usage: ".",
			Value: "gorm.io/gorm",
		},
		{
			Usage: "",
			Value: "weicai.zhao.io/tools",
		},
		{
			Usage: "",
			Value: "weicai.zhao.io/typex",
		},
	}

	t.Log(c.ToString())

	c.Consts = Consts{
		{
			Name:  "key",
			Value: `"val"`,
		},
		{
			Name:  "key1",
			Value: `"val1"`,
		},
	}
	t.Log(c.ToString())

	c.Vars = Vars{
		{
			Name:  "key",
			Value: `"val"`,
		},
		{
			Name:  "key1",
			Value: `"val1"`,
		},
	}
	t.Log(c.ToString())

	c.Structs = Structs{
		{
			StructName: "Testing",
			Usage:      "t",
			Comment:    "testing func",
			Fields: Fields{{
				FieldName: "F1",
				Type:      "string",
				Tags:      Tags{JsonTag{Value: []string{"f1"}}, GormTag{Value: []string{"f2"}}},
				Comment:   "fi field",
			},
			},
			Functions: Functions{
				{
					Imports:      nil,
					FunctionName: "T2",
					Params:       nil,
					Owner:        Owner{},
					Returns:      nil,
					Value:        "",
				},
			},
		},
		{
			StructName: "Testing",
			Usage:      "t",
			Comment:    "testing func",
			Fields: Fields{{
				FieldName: "F1",
				Type:      "string",
				Tags:      Tags{JsonTag{Value: []string{"f1"}}, GormTag{Value: []string{"f2"}}},
				Comment:   "fi field",
			},
			},
			Functions: Functions{
				{
					Imports:      nil,
					FunctionName: "T2",
					Params:       nil,
					Owner:        Owner{},
					Returns:      nil,
					Value:        "",
				},
			},
		},
	}
	t.Log(c.ToString())

	c.Functions = Functions{
		{
			Imports:      Imports{},
			FunctionName: "testName1",
			Params: []Param{
				{
					Names: []string{"t1", "t2"},
					Type:  "string",
				},
				{
					Names: []string{"t3", "t4"},
					Type:  "int64",
				},
			},
			Owner: Owner{},
			Returns: Returns{
				{
					Names: []string{},
					Type:  "string",
				},
				{
					Names: nil,
					Type:  "int64",
				},
			},
			Value: "    return",
		},
		{
			Imports:      Imports{},
			FunctionName: "testName2",
			Params: []Param{
				{
					Names: []string{"t1", "t2"},
					Type:  "string",
				},
				{
					Names: []string{"t3", "t4"},
					Type:  "int64",
				},
			},
			Owner: Owner{},
			Returns: Returns{
				{
					Names: []string{},
					Type:  "string",
				},
				{
					Names: nil,
					Type:  "int64",
				},
			},
			Value: "    return",
		},
	}
	t.Log(c.ToString())
}
