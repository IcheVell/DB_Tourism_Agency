package models

type FinancialCategory struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name          string `gorm:"column:name;not null;unique" json:"name"`
	OperationType string `gorm:"column:operation_type;not null" json:"operation_type"`
}

func (FinancialCategory) TableName() string {
	return "financial_categories"
}
