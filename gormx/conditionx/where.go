package conditionx

// Where gorm where 条件
type Where interface {
	Condition() string
	Args() []interface{}
}
