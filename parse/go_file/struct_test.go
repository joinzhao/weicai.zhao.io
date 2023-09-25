package go_file

import "testing"

func TestStruct_GetImports(t *testing.T) {
	var s = Struct{
		Fields:    Fields{{Import: Import{Value: "1"}}, {Import: Import{Value: "2"}}, {Import: Import{Value: "3"}}, {Import: Import{Value: ""}}},
		Functions: Functions{{Imports: Imports{}}, {Imports: Imports{{Value: "1"}, {Value: "2"}}}, {Imports: Imports{}}},
	}
	t.Log(s.GetImports())
}

func TestStruct_ToString(t *testing.T) {
	var s = Struct{
		StructName: "Testing",
		Usage:      "t",
		Comment:    "testing func",
	}
	t.Log(s.ToString())

	s.Fields = Fields{{
		FieldName: "F1",
		Type:      "string",
		Tags:      Tags{JsonTag{Value: []string{"f1"}}, GormTag{Value: []string{"f2"}}},
		Comment:   "fi field",
	}, {
		FieldName: "F2",
		Type:      "string",
		Tags:      Tags{},
		Comment:   "fi field",
	}, {
		FieldName: "F1",
		Type:      "string",
		Tags:      nil,
		Comment:   "",
	}}
	t.Log(s.ToString())

	s.Functions = Functions{{
		Imports:      nil,
		FunctionName: "T1",
		Params:       nil,
		Owner:        Owner{},
		Returns:      nil,
		Value:        "",
	}, {
		Imports:      nil,
		FunctionName: "T2",
		Params:       nil,
		Owner:        Owner{},
		Returns:      nil,
		Value:        "",
	}, {
		Imports:      nil,
		FunctionName: "T3",
		Params:       nil,
		Owner:        Owner{},
		Returns:      nil,
		Value:        "",
	}}
	t.Log(s.ToString())
}

func TestStructs_GetImports(t *testing.T) {
	var s = Structs{
		{
			Fields:    Fields{{Import: Import{Value: "1"}}, {Import: Import{Value: "2"}}, {Import: Import{Value: "3"}}, {Import: Import{Value: ""}}},
			Functions: Functions{{Imports: Imports{}}, {Imports: Imports{{Value: "1"}, {Value: "2"}}}, {Imports: Imports{}}},
		},
		{
			Fields:    Fields{{Import: Import{Value: "1"}}, {Import: Import{Value: "2"}}, {Import: Import{Value: "3"}}, {Import: Import{Value: ""}}},
			Functions: Functions{{Imports: Imports{}}, {Imports: Imports{{Value: "1"}, {Value: "2"}}}, {Imports: Imports{}}},
		},
		{
			Fields:    Fields{{Import: Import{Value: "1"}}, {Import: Import{Value: "2"}}, {Import: Import{Value: "3"}}, {Import: Import{Value: ""}}},
			Functions: Functions{{Imports: Imports{}}, {Imports: Imports{{Value: "1"}, {Value: "2"}}}, {Imports: Imports{}}},
		},
	}
	t.Log(s.GetImports())
}

func TestStructs_ToString(t *testing.T) {
	var s = Structs{
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

	t.Log(s.ToString())
}
