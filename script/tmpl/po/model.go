package po

const (
	ModelTableName = "model"
)

type Model struct {
	Id int64 `json:"id"`
}

func (Model) TableName() string {
	return ModelTableName
}
