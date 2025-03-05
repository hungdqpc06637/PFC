package types

type PFC_StitchingInstruction struct {
	StitchingInstructionID string `gorm:"column:StitchingInstructionID"`
	ModelType              string `gorm:"column:ModelType"`
	ModelName              string `gorm:"column:ModelName"`
	MaterialNumber         string `gorm:"column:MaterialNumber"`
	Title                  string `gorm:"column:Title"`
	ItemIndex              string `gorm:"column:ItemIndex"`
}

type PFC_ItemStitchingInstruction struct {
	ItemStitchingInstructionID string `gorm:"column:ItemCCCID"`
	StitchingInstructionID     string `gorm:"column:StitchingInstructionID"`
	Component                  string `gorm:"column:Component"`
	ImageContent               string `gorm:"column:ImageContent"`
	McType                     string `gorm:"column:McType"`
	NeedleSystem               string `gorm:"column:NeedleSystem"`
	NeedleSize                 string `gorm:"column:NeedleSize"`
	NeedlePointType            string `gorm:"column:NeedlePointType"`
	ThreadType                 string `gorm:"column:ThreadType"`
	StitchingMargin            string `gorm:"column:StitchingMargin"`
	StitchPerInch              string `gorm:"column:StitchPerInch"`
	AttachingMethod            string `gorm:"column:AttachingMethod"`
	StitchingGuideName         string `gorm:"column:StitchingGuideName"`
	ItemIndex                  string `gorm:"column:ItemIndex"`
}
