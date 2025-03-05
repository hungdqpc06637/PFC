package types

type PFC_BottomSilkScreenProcess struct {
	BottomSilkScreenProcessID string `gorm:"column:BottomSilkScreenProcessID"`
	ModelType                string `gorm:"column:ModelType"`
	ModelName                string `gorm:"column:ModelName"`
	MaterialNumber           string `gorm:"column:MaterialNumber"`
	Title                    string `gorm:"column:Title"`
	ItemIndex                string `gorm:"column:ItemIndex"`
}

type PFC_ItemBottomSilkScreenProcess struct {
	ItemBottomSilkScreenProcessID string `gorm:"column:ItemBottomSilkScreenProcessID"`
	BottomSilkScreenProcessID     string `gorm:"column:BottomSilkScreenProcessID"`
	ComponentName                     string `gorm:"column:ComponentName"`
	Material                       string `gorm:"column:Material"`
	Vendor                      string `gorm:"column:Vendor"`
	TableRow1                      string `gorm:"column:TableRow1"`
	Remarks                          string `gorm:"column:Remarks"`
	Size                          string `gorm:"column:Size"`
	TotalWBSB                          string `gorm:"column:TotalWBSB"`
}
