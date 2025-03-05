package types

type PFC_UpperLogoSpecification struct {
	UpperLogoSpecificationID string `gorm:"column:UpperLogoSpecificationID"`
	ModelType                string `gorm:"column:ModelType"`
	ModelName                string `gorm:"column:ModelName"`
	MaterialNumber           string `gorm:"column:MaterialNumber"`
	Title                    string `gorm:"column:Title"`
	ItemIndex                string `gorm:"column:ItemIndex"`
}

type PFC_ItemUpperLogoSpecification struct {
	ItemUpperLogoSpecificationID string `gorm:"column:ItemUpperLogoSpecificationID"`
	UpperLogoSpecificationID     string `gorm:"column:UpperLogoSpecificationID"`
	Component                    string `gorm:"column:Component"`
	ImageContent                 string `gorm:"column:ImageContent"`
	Vendor                       string `gorm:"column:Vendor"`
	Material                     string `gorm:"column:Material"`
	TableRow1                    string `gorm:"column:TableRow1"`
	TableRow2                    string `gorm:"column:TableRow2"`
	TableRow3                    string `gorm:"column:TableRow3"`
	TableRow4                    string `gorm:"column:TableRow4"`
	TableRow5                    string `gorm:"column:TableRow5"`
	TableRow6                    string `gorm:"column:TableRow6"`
	ItemIndex                    string `gorm:"column:ItemIndex"`
}
