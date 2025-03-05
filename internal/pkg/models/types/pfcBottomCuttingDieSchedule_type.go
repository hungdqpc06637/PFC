package types

type PFC_BottomCuttingDieSchedule struct {
	BottomCuttingDieScheduleID string `gorm:"column:BottomCuttingDieScheduleID"`
	ModelType                  string `gorm:"column:ModelType"`
	ModelName                  string `gorm:"column:ModelName"`
	MaterialNumber             string `gorm:"column:MaterialNumber"`
	Title                      string `gorm:"column:Title"`
	ItemIndex                  string `gorm:"column:ItemIndex"`
}

type PFC_ItemBottomCuttingDieSchedule struct {
	ItemBottomCuttingDieScheduleID string `gorm:"column:ItemBottomCuttingDieScheduleID"`
	BottomCuttingDieScheduleID     string `gorm:"column:BottomCuttingDieScheduleID"`
	Component                      string `gorm:"column:Component"`
	ImageContent                   string `gorm:"column:ImageContent"`
	SizeRange1                     string `gorm:"column:SizeRange1"`
	SizeRange2                     string `gorm:"column:SizeRange2"`
	SizeRange                      string `gorm:"column:SizeRange"`
	SizeRangeAreSame               string `gorm:"column:SizeRangeAreSame"`
	Remarks                        string `gorm:"column:Remarks"`
	NumberOfLayers                 string `gorm:"column:NumberOfLayers"`
	Thickness                      string `gorm:"column:Thickness"`
	Width                          string `gorm:"column:Width"`
	Hardness                       string `gorm:"column:Hardness"`
	ItemIndex                      string `gorm:"column:ItemIndex"`
}
