package types

type PFC_OutsideConveyorProcess struct {
	OutsideConveyorProcessID string `gorm:"column:OutsideConveyorProcessID"`
	ModelType                string `gorm:"column:ModelType"`
	ModelName                string `gorm:"column:ModelName"`
	MaterialNumber           string `gorm:"column:MaterialNumber"`
	Title                    string `gorm:"column:Title"`
	ItemIndex                string `gorm:"column:ItemIndex"`
}

type PFC_ItemOutsideConveyorProcess struct {
	ItemOutsideConveyorProcessID string `gorm:"column:ItemOutsideConveyorProcessID"`
	OutsideConveyorProcessID     string `gorm:"column:OutsideConveyorProcessID"`
	ComponentName                string `gorm:"column:ComponentName"`
	Material                     string `gorm:"column:Material"`
	Vendor                       string `gorm:"column:Vendor"`
	TableRow1                    string `gorm:"column:TableRow1"`
}
