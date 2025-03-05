package types

type PFC_SecondProcess struct {
	SecondProcessID     string                  `gorm:"column:SecondProcessID"`
	ModelType           string                  `gorm:"column:ModelType"`
	ModelName           string                  `gorm:"column:ModelName"`
	MaterialNumber      string                  `gorm:"column:MaterialNumber"`
	Title               string                  `gorm:"column:Title"`
	ItemIndex           string                  `gorm:"column:ItemIndex"`
}

type PFC_ItemSecondProcess struct {
	ItemSecondProcessID string `gorm:"column:ItemSecondProcessID"`
	SecondProcessID     string `gorm:"column:SecondProcessID"`
	Component           string `gorm:"column:Component"`
	ImageContent        string `gorm:"column:ImageContent"`
	Material            string `gorm:"column:Material"`
	Method              string `gorm:"column:Method"`
	ItemIndex           string `gorm:"column:ItemIndex"`
}
