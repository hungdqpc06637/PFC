package types

type PFC_BottomPartsProcess struct {
	BottomPartsProcessID string `gorm:"column:BottomPartsProcessID"`
	ModelType            string `gorm:"column:ModelType"`
	ModelName            string `gorm:"column:ModelName"`
	MaterialNumber       string `gorm:"column:MaterialNumber"`
	Title                string `gorm:"column:Title"`
	ItemIndex            string `gorm:"column:ItemIndex"`
}

type PFC_ItemBottomPartsProcess struct {
	ItemBottomPartsProcessID string `gorm:"column:ItemBottomPartsProcessID"`
	BottomPartsProcessID     string `gorm:"column:BottomPartsProcessID"`
	Component                string `gorm:"column:Component"`
	Material                 string `gorm:"column:Material"`
	Vendor                   string `gorm:"column:Vendor"`
	TableRow1                string `gorm:"column:TableRow1"`
	RemarksImages            string `gorm:"column:RemarksImages"`
	RemarksSize              string `gorm:"column:RemarksSize"`
}
