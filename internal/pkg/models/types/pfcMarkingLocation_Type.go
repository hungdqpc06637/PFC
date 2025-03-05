package types

type PFC_MarkingLocation struct {
	MarkingLocationID     string                    `gorm:"column:MarkingLocationID"`
	ModelType             string                    `gorm:"column:ModelType"`
	ModelName             string                    `gorm:"column:ModelName"`
	MaterialNumber        string                    `gorm:"column:MaterialNumber"`
	Title                 string                    `gorm:"column:Title"`
	ItemIndex             string                    `gorm:"column:ItemIndex"`
}

type PFC_ItemMarkingLocation struct {
	ItemMarkingLocationID string `gorm:"column:ItemMarkingLocationID"`
	MarkingLocationID     string `gorm:"column:MarkingLocationID"`
	Component             string `gorm:"column:Component"`
	ImageContent          string `gorm:"column:ImageContent"`
	TitleImage            string `gorm:"column:TitleImage"`
	Process               string `gorm:"column:Process"`
	ItemIndex             string `gorm:"column:ItemIndex"`
}
