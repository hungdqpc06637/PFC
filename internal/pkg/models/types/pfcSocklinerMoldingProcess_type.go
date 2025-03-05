package types

type PFC_SocklinerMoldingProcess struct {
	SocklinerMoldingProcessID string `gorm:"column:SocklinerMoldingProcessID"`
	ModelType                 string `gorm:"column:ModelType"`
	ModelName                 string `gorm:"column:ModelName"`
	MaterialNumber            string `gorm:"column:MaterialNumber"`
	Title                     string `gorm:"column:Title"`
	ItemIndex                 string `gorm:"column:ItemIndex"`
}

type PFC_ItemSocklinerMoldingProcess struct {
	ItemSocklinerMoldingProcessID string `gorm:"column:ItemSocklinerMoldingProcessID"`
	SocklinerMoldingProcessID     string `gorm:"column:SocklinerMoldingProcessID"`
	RawMaterialName               string `gorm:"column:ComponentName"`
	Vendor                        string `gorm:"column:Vendor"`
	TableRow1                     string `gorm:"column:TableRow1"`
	Remarks                       string `gorm:"column:Remarks"`
	RemarksImage                  string `gorm:"column:RemarksImage"`
	Model                         string `gorm:"column:Model"`
	Size                          string `gorm:"column:Size"`
}
