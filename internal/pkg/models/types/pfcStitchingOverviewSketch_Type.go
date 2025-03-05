package types

type PFC_StitchingOverviewSketch struct {
	StitchingOverviewSketchID     string                            `gorm:"column:StitchingOverviewSketchID"`
	ModelType                     string                            `gorm:"column:ModelType"`
	ModelName                     string                            `gorm:"column:ModelName"`
	MaterialNumber                string                            `gorm:"column:MaterialNumber"`
	Title                         string                            `gorm:"column:Title"`
	ItemIndex                     string                            `gorm:"column:ItemIndex"`
}

type PFC_ItemStitchingOverviewSketch struct {
	ItemStitchingOverviewSketchID string `gorm:"column:ItemStitchingOverviewSketchID"`
	StitchingOverviewSketchID     string `gorm:"column:StitchingOverviewSketchID"`
	Component                     string `gorm:"column:Component"`
	ImageContent                  string `gorm:"column:ImageContent"`
	RightFoot                     string `gorm:"column:RightFoot"`
	ItemIndex                     string `gorm:"column:ItemIndex"`
}
