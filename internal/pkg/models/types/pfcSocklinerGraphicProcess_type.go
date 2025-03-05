package types

type PFC_SocklinerGraphicProcess struct {
	SocklinerGraphicProcessID string `gorm:"column:SocklinerGraphicProcessID"`
	ModelType                 string `gorm:"column:ModelType"`
	ModelName                 string `gorm:"column:ModelName"`
	MaterialNumber            string `gorm:"column:MaterialNumber"`
	Title                     string `gorm:"column:Title"`
	ItemIndex                 string `gorm:"column:ItemIndex"`
}

type PFC_ItemSocklinerGraphicProcess struct {
	ItemSocklinerGraphicProcessID string `gorm:"column:ItemSocklinerGraphicProcessID"`
	SocklinerGraphicProcessID     string `gorm:"column:SocklinerGraphicProcessID"`
	ComponentName                 string `gorm:"column:ComponentName"`
	SocklinerLogo                 string `gorm:"column:SocklinerLogo"`
	Vendor                        string `gorm:"column:Vendor"`
	TableRow1                     string `gorm:"column:TableRow1"`
	Remarks                       string `gorm:"column:Remarks"`
	RemarksImage                  string `gorm:"column:RemarksImage"`
	ModelSize                     string `gorm:"column:ModelSize"`
}
