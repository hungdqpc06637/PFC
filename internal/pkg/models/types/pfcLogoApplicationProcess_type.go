package types

type PFC_LogoApplicationProcess struct {
	LogoApplicationProcessID string `gorm:"column:LogoApplicationProcessID"`
	ModelType                string `gorm:"column:ModelType"`
	ModelName                string `gorm:"column:ModelName"`
	MaterialNumber           string `gorm:"column:MaterialNumber"`
	Title                    string `gorm:"column:Title"`
	ItemIndex                string `gorm:"column:ItemIndex"`
}

type PFC_ItemLogoApplicationProcess struct {
	ItemLogoApplicationProcessID string `gorm:"column:ItemLogoApplicationProcessID"`
	LogoApplicationProcessID     string `gorm:"column:LogoApplicationProcessID"`
	ComponentName                string `gorm:"column:ComponentName"`
	LogoSockliner                string `gorm:"column:LogoSockliner"`
	Vendor                       string `gorm:"column:Vendor"`
	TableRow1                    string `gorm:"column:TableRow1"`
	Remarks                      string `gorm:"column:Remarks"`
	RemarksImage                 string `gorm:"column:RemarksImage"`
	Size                         string `gorm:"column:Size"`
}
