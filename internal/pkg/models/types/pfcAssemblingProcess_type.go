package types

type PFC_AssemblingProcess struct {
	AssemblingProcessID string `gorm:"column:AssemblingProcessID"`
	ModelType           string `gorm:"column:ModelType"`
	ModelName           string `gorm:"column:ModelName"`
	MaterialNumber      string `gorm:"column:MaterialNumber"`
	Title               string `gorm:"column:Title"`
	ItemIndex           string `gorm:"column:ItemIndex"`
}

type PFC_ItemAssemblingProcess struct {
	ItemAssemblingProcessID  string `gorm:"column:ItemAssemblingProcessID"`
	AssemblingProcessID      string `gorm:"column:AssemblingProcessID"`
	TableRow1                string `gorm:"column:TableRow1"`
	ImagesContent            string `gorm:"column:ImagesContent"`
	LastingLaceType          string `gorm:"column:LastingLaceType"`
	LastingClipSchedule      string `gorm:"column:LastingClipSchedule"`
	ChillerMoldSpecification string `gorm:"column:ChillerMoldSpecification"`
	TotalWBSB                string `gorm:"column:TotalWBSB"`
}
