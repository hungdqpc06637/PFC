package types

type PFC_RubberComponentSpecification struct {
	RubberComponentSpecificationID string `gorm:"column:RubberComponentSpecificationID"`
	ModelType                      string `gorm:"column:ModelType"`
	ModelName                      string `gorm:"column:ModelName"`
	MaterialNumber                 string `gorm:"column:MaterialNumber"`
	Title                          string `gorm:"column:Title"`
	ItemIndex                      string `gorm:"column:ItemIndex"`
}

type PFC_ItemRubberComponentSpecification struct {
	ItemRubberComponentSpecificationID string `gorm:"column:ItemRubberComponentSpecificationID"`
	RubberComponentSpecificationID     string `gorm:"column:RubberComponentSpecificationID"`
	TableRow1                          string `gorm:"column:TableRow1"`
	PartInformation                    string `gorm:"column:PartInformation"`
	LengthGrading                      string `gorm:"column:LengthGrading"`
	TableRow2                          string `gorm:"column:TableRow2"`
}
