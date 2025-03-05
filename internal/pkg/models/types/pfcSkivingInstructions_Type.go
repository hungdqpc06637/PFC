package types

type PFC_SkivingInstructions struct {
	SkivingInstructionsID     string                        `gorm:"column:SkivingInstructionsID"`
	ModelType                 string                        `gorm:"column:ModelType"`
	ModelName                 string                        `gorm:"column:ModelName"`
	MaterialNumber            string                        `gorm:"column:MaterialNumber"`
	Title                     string                        `gorm:"column:Title"`
	SkivingKey                string                        `gorm:"column:SkivingKey"`
	ItemIndex                 string                        `gorm:"column:ItemIndex"`
}

type PFC_ItemSkivingInstructions struct {
	ItemSkivingInstructionsID string `gorm:"column:ItemSkivingInstructionsID"`
	SkivingInstructionsID     string `gorm:"column:SkivingInstructionsID"`
	Component                 string `gorm:"column:Component"`
	ImageContent              string `gorm:"column:ImageContent"`
	SkivedEdgeThickness       string `gorm:"column:SkivedEdgeThickness"`
	SkivingWidth              string `gorm:"column:SkivingWidth"`
	ItemIndex                 string `gorm:"column:ItemIndex"`
}
