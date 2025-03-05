package types

type PFC_PressingPadSpecification struct {
	PressingPadSpecificationID string `gorm:"column:PressingPadSpecificationID"`
	ModelType                  string `gorm:"column:ModelType"`
	ModelName                  string `gorm:"column:ModelName"`
	MaterialNumber             string `gorm:"column:MaterialNumber"`

	ColorWayID  string `gorm:"column:ColorWayID"`
	TechLevel   string `gorm:"column:TechLevel"`
	FirstSource string `gorm:"column:FirstSource"`

	Title     string `gorm:"column:Title"`
	ItemIndex string `gorm:"column:ItemIndex"`
}

type PFC_ItemPressingPadSpecification struct {
	ItemPressingPadSpecificationID string `gorm:"column:ItemPressingPadSpecificationID"`
	PressingPadSpecificationID     string `gorm:"column:PressingPadSpecificationID"`
	TableRow1                      string `gorm:"column:TableRow1"`
	TableRow2                      string `gorm:"column:TableRow2"`
}
