package types

type PFC_PerforationSpecification struct {
	PerforationSpecificationID string `gorm:"column:PerforationSpecificationID"`
	ModelType                  string `gorm:"column:ModelType"`
	ModelName                  string `gorm:"column:ModelName"`
	MaterialNumber             string `gorm:"column:MaterialNumber"`
	Title                      string `gorm:"column:Title"`
	ItemIndex                  string `gorm:"column:ItemIndex"`
}

type PFC_ItemPerforationSpecification struct {
	ItemPerforationSpecificationID string `gorm:"column:ItemPerforationSpecificationID"`
	PerforationSpecificationID     string `gorm:"column:PerforationSpecificationID"`
	Component                      string `gorm:"column:Component"`
	ImageContent                   string `gorm:"column:ImageContent"`
	SizeGroup1                     string `gorm:"column:SizeGroup1"`
	SizeGroup2                     string `gorm:"column:SizeGroup2"`
	SizeGroup3                     string `gorm:"column:SizeGroup3"`
	ItemIndex                      string `gorm:"column:ItemIndex"`
}
