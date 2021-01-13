package domains

func (Type) TableName() string {
	return "type"
}

type Type struct {
	ID   int32  `gorm:"column:id; PRIMARY_KEY" json:"id"`
	Type string `gorm:"column:type;" json:"type"`
}
