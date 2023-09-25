package go_file

import "testing"

func TestField_GetImport(t *testing.T) {

}

func TestField_ToString(t *testing.T) {
	var f = Field{
		Import:    Import{},
		FieldName: "Field",
		Type:      "int64",
		Tags: []Tag{
			JsonTag{
				Key:   "",
				Value: []string{"field"},
			},
			GormTag{
				Key:   "",
				Value: []string{"column:field"},
			},
		},
		Comment: "field test",
	}

	t.Log(f.ToString())

	f.Comment = ""
	t.Log(f.ToString())

	f.Tags = nil
	t.Log(f.ToString())
}

func TestFields_GetImports(t *testing.T) {
	var f = Fields{}
	t.Log(f.GetImports())

	f = append(f, Field{
		Import: Import{},
	})
	t.Log(f.GetImports())

	f = append(f, Field{
		Import: Import{Value: "t1"},
	})
	t.Log(f.GetImports())

	f = append(f, Field{
		Import: Import{Value: "t1"},
	})
	t.Log(f.GetImports())

	f = append(f, Field{
		Import: Import{Value: "t12"},
	})
	t.Log(f.GetImports())
}

func TestFields_ToString(t *testing.T) {
	var f = Fields{
		{
			Import:    Import{Value: ""},
			FieldName: "Field",
			Type:      "int64",
			Tags: []Tag{
				JsonTag{
					Key:   "",
					Value: []string{"field"},
				},
				GormTag{
					Key:   "",
					Value: []string{"column:field"},
				},
			},
			Comment: "field test",
		},
		{
			Import:    Import{Value: ""},
			FieldName: "Field1",
			Type:      "int64",
			Tags:      []Tag{},
			Comment:   "field test",
		},
		{
			Import:    Import{Value: ""},
			FieldName: "Field",
			Type:      "int64",
			Tags: []Tag{
				JsonTag{
					Key:   "",
					Value: []string{"field"},
				},
				GormTag{
					Key:   "",
					Value: []string{"column:field"},
				},
			},
			Comment: "",
		},
		{
			Import:    Import{Value: ""},
			FieldName: "Field",
			Type:      "int64",
			Tags:      []Tag{},
			Comment:   "",
		},
		{
			Import:    Import{Value: ""},
			FieldName: "Field",
			Type:      "int64",
			Tags:      []Tag{},
			Comment:   "",
		},
	}

	t.Log(f.ToString())

	f = Fields{{
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
	},
	}
	t.Log(f.ToString())
}
