package types

type PFC_OutsoleSpecification struct {
	OutsoleSpecificationID string `gorm:"column:OutsoleSpecificationID"`
	ModelType              string `gorm:"column:ModelType"`
	ModelName              string `gorm:"column:ModelName"`
	MaterialNumber         string `gorm:"column:MaterialNumber"`
	Title                  string `gorm:"column:Title"`
	ItemIndex              string `gorm:"column:ItemIndex"`
}

type PFC_ItemOutsoleSpecification struct {
	ItemOutsoleSpecificationID string `gorm:"column:ItemOutsoleSpecificationID"`
	OutsoleSpecificationID     string `gorm:"column:OutsoleSpecificationID"`
	TableRow1                  string `gorm:"column:TableRow1"`
	PreformPartsNestingDiagram string `gorm:"column:PreformPartsNestingDiagram"`
	FinalPart                  string `gorm:"column:FinalPart"`
	TableRow2                  string `gorm:"column:TableRow2"`
	TotalWeightPerSize         string `gorm:"column:TotalWeightPerSize"`
}
