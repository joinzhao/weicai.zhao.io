package reportx

import (
	"weicai.zhao.io/gormx/conditionx"
	"weicai.zhao.io/gormx/tablex"
)

// Report 报告
type Report interface {
	ReportReader
	ReportWriter
}

// ReportReader 只读报告
type ReportReader interface {
	First(tablex.Tabler, ...conditionx.Where) error
}

// ReportWriter 只写报告
type ReportWriter interface {
	CreateInBatches(tablers ...tablex.Tabler) error
	Create(tablex.Tabler) error
	Update(tablex.Tabler, ...conditionx.Where) error
	Delete(tablex.Tabler, ...conditionx.Where) error
}
