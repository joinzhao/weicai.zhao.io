package option

import "gorm.io/gorm"

type Where struct {
	Column string
	Expr   conditionExpression
	Value  []interface{}
}

func (op Where) apply(db *gorm.DB) *gorm.DB {
	return db.Where(op.Expr.toQuery(op.Column), op.Value...)
}

type Or struct {
	Column string
	Expr   conditionExpression
	Value  []interface{}
}

func (op Or) apply(db *gorm.DB) *gorm.DB {
	return db.Or(op.Expr.toQuery(op.Column), op.Value...)
}
