package types

type PFC_OutsolePressingProcess struct {
	OutsolePressingProcessID string `gorm:"column:OutsolePressingProcessID"`
	ModelType                string `gorm:"column:ModelType"`
	ModelName                string `gorm:"column:ModelName"`
	MaterialNumber           string `gorm:"column:MaterialNumber"`
	Title                    string `gorm:"column:Title"`
	ItemIndex                string `gorm:"column:ItemIndex"`
}

type PFC_ItemOutsolePressingProcess struct {
	ItemOutsolePressingProcessID string `gorm:"column:ItemOutsolePressingProcessID"`
	OutsolePressingProcessID     string `gorm:"column:OutsolePressingProcessID"`

	TableRow1 string `gorm:"column:TableRow1"`
	TableRow2 string `gorm:"column:TableRow2"`

	AverageWastePerPart    string `gorm:"column:AverageWastePerPart"`
	VerageWasteRatePerPart string `gorm:"column:VerageWasteRatePerPart"`
}
