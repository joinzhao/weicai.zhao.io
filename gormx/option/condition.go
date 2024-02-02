package option

type condition string

const (
	EqualCondition            condition = "="
	NotEqualCondition         condition = "<>"
	InCondition               condition = "IN"
	LikeCondition             condition = "LIKE"
	LessThanCondition         condition = "<"
	GreaterThanCondition      condition = ">"
	LessThanEqualCondition    condition = "<="
	GreaterThanEqualCondition condition = ">="
)

var (
	Equal            = singleCondition{cond: EqualCondition}
	NotEqual         = singleCondition{cond: NotEqualCondition}
	Like             = singleCondition{cond: LikeCondition}
	LessThan         = singleCondition{cond: LessThanCondition}
	GreaterThan      = singleCondition{cond: GreaterThanCondition}
	LessThanEqual    = singleCondition{cond: LessThanEqualCondition}
	GreaterThanEqual = singleCondition{cond: GreaterThanEqualCondition}
	In               = inCondition{}
)
