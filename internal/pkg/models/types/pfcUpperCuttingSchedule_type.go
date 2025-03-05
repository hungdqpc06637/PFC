package types

type PFC_UpperCuttingSchedule struct {
	UpperCuttingDieScheduleID      string                         `gorm:"column:UpperCuttingDieScheduleID"`
	ModelType                      string                         `gorm:"column:ModelType"`
	ModelName                      string                         `gorm:"column:ModelName"`
	MaterialNumber                 string                         `gorm:"column:MaterialNumber"`
	ComponentName                  string                         `gorm:"column:ComponentName"`
	ImageRemark                    string                         `gorm:"column:ImageRemark"`
	Remark                         string                         `gorm:"column:Remark"`
	PageIndex                      string                         `gorm:"column:PageIndex"`
	ItemsUpperCuttingDieScheduleID []PFC_ItemUpperCuttingSchedule `gorm:"-" json:"RoItemUpperCuttingSchedulells"`
}

type PFC_ItemUpperCuttingSchedule struct {
	ItemUpperCuttingDieScheduleID string `gorm:"column:ItemUpperCuttingDieScheduleID"`
	UpperCuttingDieScheduleID     string `gorm:"column:UpperCuttingDieScheduleID"`
	ComponentName                 string `gorm:"column:ComponentName"`
	ItemIndex                     string `gorm:"column:ItemIndex"`
	SizeRange1                    string `gorm:"column:SizeRange1"`
	SizeRange2                    string `gorm:"column:SizeRange2"`
	SizeRange                     string `gorm:"column:SizeRange"`
	SizeRangeAreSame              string `gorm:"column:SizeRangeAreSame"`
	ImageContent                  string `gorm:"column:ImageContent"`
	ImageRemark                   string `gorm:"column:ImageRemark"`
	NumberOfPlayers               string `gorm:"column:NumberOfPlayers"`
}
