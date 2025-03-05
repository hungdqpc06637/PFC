package types

type PFC_ComputerStitchingSchedule struct {
	ComputerStitchingScheduleID     string                              `gorm:"column:ComputerStitchingScheduleID"`
	ModelType                       string                              `gorm:"column:ModelType"`
	ModelName                       string                              `gorm:"column:ModelName"`
	MaterialNumber                  string                              `gorm:"column:MaterialNumber"`
	Title                           string                              `gorm:"column:Title"`
	ItemIndex                       string                              `gorm:"column:ItemIndex"`
}

type PFC_ItemComputerStitchingSchedule struct {
	ItemComputerStitchingScheduleID string `gorm:"column:ItemComputerStitchingScheduleID"`
	ComputerStitchingScheduleID     string `gorm:"column:ComputerStitchingScheduleID"`
	Component                       string `gorm:"column:Component"`
	ImageContent                    string `gorm:"column:ImageContent"`
	StitchingMargin                 string `gorm:"column:StitchingMargin"`
	NeedleTypeSize                  string `gorm:"column:NeedleTypeSize"`
	StitchPerInch                   string `gorm:"column:StitchPerInch	"`
	Size                            string `gorm:"column:Size"`
	ItemIndex                       string `gorm:"column:ItemIndex"`
}
