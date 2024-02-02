package option

import "fmt"

type conditionExpression interface {
	toQuery(column string) string
}

type (
	singleCondition struct {
		cond condition
	}
	inCondition struct {
	}
)

func (c singleCondition) toQuery(column string) string {
	return fmt.Sprintf("`%s` %s ?", column, c.cond)
}
func (inCondition) toQuery(column string) string {
	return fmt.Sprintf("`%s` %s (?)", column, InCondition)
}
