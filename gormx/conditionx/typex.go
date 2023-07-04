package conditionx

import (
	"fmt"
	"weicai.zhao.io/gormx/tablex"
)

type Id int64

func (id Id) Condition() string {
	return "id = ?"
}
func (id Id) Args() []interface{} {
	return []interface{}{id}
}

type Uuid string

func (id Uuid) Condition() string {
	return "uuid = ?"
}
func (id Uuid) Args() []interface{} {
	return []interface{}{id}
}

type Status uint8

func (id Status) Condition() string {
	return "status = ?"
}
func (id Status) Args() []interface{} {
	return []interface{}{id}
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
