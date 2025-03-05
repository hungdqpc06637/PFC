package types

type PFC_PageSetup struct {
	PageSetupID          string `gorm:"column:PageSetupID"`
	ModelType            string `gorm:"column:ModelType"`
	ModelName            string `gorm:"column:ModelName"`
	MaterialNumber       string `gorm:"column:MaterialNumber"`
	LeftSelectionHeader  string `gorm:"column:LeftSelectionHeader"`
	RightSelectionHeader string `gorm:"column:RightSelectionHeader"`
}
