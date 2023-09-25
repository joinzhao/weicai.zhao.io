package go_file

import "testing"

func TestFunction_ToString(t *testing.T) {
	var f = Function{
		Imports:      Imports{},
		FunctionName: "testName",
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
	}

	t.Log(">\n", f.ToString(), "\n<")
}

func TestFunction_GetImports(t *testing.T) {
	var f = Function{
		Imports: nil,
	}
	t.Log(f.GetImports())

	f.Imports = Imports{}
	t.Log(f.GetImports())

	f.Imports = Imports{
		Import{Value: "t"},
		Import{Value: "t"},
		Import{Value: "t1"},
	}
	t.Log(f.GetImports())
}

func TestFunctions_ToString(t *testing.T) {
	var f = Functions{
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

	t.Log(">\n", f.ToString(), "\n<")
}

func TestFunctions_GetImports(t *testing.T) {
	var f = Functions{
		{
			Imports: nil,
		},
		{
			Imports: Imports{{Value: ""}},
		},
		{
			Imports: Imports{{Value: "t1"}},
		},
		{
			Imports: Imports{{Value: "t1"}, {Value: "t2"}},
		},
	}
	t.Log(f.GetImports())
}
