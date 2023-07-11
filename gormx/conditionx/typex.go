package conditionx

import (
	"fmt"
	"weicai.zhao.io/gormx/tablex"
)

type Id int64

func (c Id) Condition() string {
	return "id = ?"
}
func (c Id) Args() []interface{} {
	return []interface{}{c}
}

type Uuid string

func (c Uuid) Condition() string {
	return "uuid = ?"
}
func (c Uuid) Args() []interface{} {
	return []interface{}{c}
}

type Code string

func (c Code) Condition() string {
	return "code = ?"
}
func (c Code) Args() []interface{} {
	return []interface{}{c}
}

type Status uint8

func (c Status) Condition() string {
	return "status = ?"
}
func (c Status) Args() []interface{} {
	return []interface{}{c}
}

type Name string

func (name Name) Condition() string {
	return "name = ?"
}
func (name Name) Args() []interface{} {
	return []interface{}{name}
}

// condition with table name
type whereWithTable struct {
	condition string
	args      []interface{}
}

func (w whereWithTable) Condition() string {
	return w.condition
}
func (w whereWithTable) Args() []interface{} {
	return w.args
}

func WithTableName(tabler tablex.Tabler, where Where) Where {
	return whereWithTable{
		condition: fmt.Sprintf("`%s`.%s", tabler.TableName(), where.Condition()),
		args:      where.Args(),
	}
}
