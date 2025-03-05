package types

type PFC_KeyManufacturingDetails struct {
	KeyManufacturingDetailsID string `gorm:"column:KeyManufacturingDetailsID"`
	ModelType                 string `gorm:"column:ModelType"`
	ModelName                 string `gorm:"column:ModelName"`
	MaterialNumber            string `gorm:"column:MaterialNumber"`
	Title                     string `gorm:"column:Title"`
	ItemIndex                 string `gorm:"column:ItemIndex"`
}

type PFC_ItemKeyManufacturingDetails struct {
	ItemKeyManufacturingDetailsID string `gorm:"column:ItemKeyManufacturingDetailsID"`
	KeyManufacturingDetailsID     string `gorm:"column:KeyManufacturingDetailsID"`
	KeyManufacturingProcess       string `gorm:"column:KeyManufacturingProcess"`
	DetailPicture                 string `gorm:"column:DetailPicture"`
	ProcessDetailKeyCheckPoint    string `gorm:"column:ProcessDetailKeyCheckPoint"`
}
