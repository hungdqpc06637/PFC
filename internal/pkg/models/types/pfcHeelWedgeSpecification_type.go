package types

type PFC_HeelWedgeSpecification struct {
	HeelWedgeSpecificationID string `gorm:"column:HeelWedgeSpecificationID"`
	ModelType                string `gorm:"column:ModelType"`
	ModelName                string `gorm:"column:ModelName"`
	MaterialNumber           string `gorm:"column:MaterialNumber"`
	Title                    string `gorm:"column:Title"`
	ItemIndex                string `gorm:"column:ItemIndex"`
}

type PFC_ItemHeelWedgeSpecification struct {
	ItemHeelWedgeSpecificationID string `gorm:"column:ItemHeelWedgeSpecificationID"`
	HeelWedgeSpecificationID     string `gorm:"column:HeelWedgeSpecificationID"`
	TableRow1                    string `gorm:"column:TableRow1"`
	Thickness                    string `gorm:"column:Thickness"`
	ImagesContent                string `gorm:"column:ImagesContent"`
}
