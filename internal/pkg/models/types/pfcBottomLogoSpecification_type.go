package types

type PFC_BottomLogoSpecification struct {
	BottomLogoSpecificationID string `gorm:"column:BottomLogoSpecificationID"`
	ModelType                 string `gorm:"column:ModelType"`
	ModelName                 string `gorm:"column:ModelName"`
	MaterialNumber            string `gorm:"column:MaterialNumber"`
	Title                     string `gorm:"column:Title"`
	ItemIndex                 string `gorm:"column:ItemIndex"`
}

type PFC_ItemBottomLogoSpecification struct {
	ItemBottomLogoSpecificationID string `gorm:"column:ItemBottomLogoSpecificationID"`
	BottomLogoSpecificationID     string `gorm:"column:BottomLogoSpecificationID"`
	Component                     string `gorm:"column:Component"`
	ImageContent                  string `gorm:"column:ImageContent"`
	Vendor                        string `gorm:"column:Vendor"`
	MaterialApplication           string `gorm:"column:MaterialApplication"`
	Model                         string `gorm:"column:Model"`
	Size                          string `gorm:"column:Size"`
	ItemIndex                     string `gorm:"column:ItemIndex"`
}
