package types

type PFC_SocklinerSpecification struct {
	SocklinerSpecificationID string `gorm:"column:SocklinerSpecificationID"`
	ModelType                string `gorm:"column:ModelType"`
	ModelName                string `gorm:"column:ModelName"`
	MaterialNumber           string `gorm:"column:MaterialNumber"`
	Title                    string `gorm:"column:Title"`
	ItemIndex                string `gorm:"column:ItemIndex"`
}

type PFC_ItemSocklinerSpecification struct {
	ItemSocklinerSpecificationID string `gorm:"column:ItemSocklinerSpecificationID"`
	SocklinerSpecificationID     string `gorm:"column:SocklinerSpecificationID"`
	Front                        string `gorm:"column:Front"`
	Back                         string `gorm:"column:Back"`
	Size                         string `gorm:"column:Size"`
	TableRow1                    string `gorm:"column:TableRow1"`
	TableRow2                    string `gorm:"column:TableRow2"`
}
