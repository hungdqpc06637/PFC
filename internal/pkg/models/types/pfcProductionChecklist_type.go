package types

type PFC_ProductionChecklist struct {
	ProductionChecklistID string `gorm:"column:ProductionChecklistID"`
	ModelType             string `gorm:"column:ModelType"`
	ModelName             string `gorm:"column:ModelName"`
	MaterialNumber        string `gorm:"column:MaterialNumber"`
	Title                 string `gorm:"column:Title"`
	ItemIndex             string `gorm:"column:ItemIndex"`
}

type PFC_ItemProductionChecklist struct {
	ItemProductionChecklistID string `gorm:"column:ItemProductionChecklistID"`
	ProductionChecklistID     string `gorm:"column:ProductionChecklistID"`
	ShoeLacingInstruction     string `gorm:"column:ShoeLacingInstruction"`
	OtherAccessoryInfo        string `gorm:"column:OtherAccessoryInfo"`
	IDSMeasurement            string `gorm:"column:IDSMeasurement"`
	InnerBoxSchedule          string `gorm:"column:InnerBoxSchedule"`
	LaceSchedule              string `gorm:"column:LaceSchedule"`
	WrappingPaperSchedule     string `gorm:"column:WrappingPaperSchedule"`
	ArchCookieSchedule        string `gorm:"column:ArchCookieSchedule"`
	CardboardFootForm         string `gorm:"column:CardboardFootForm"`
	MoldedFootForm            string `gorm:"column:MoldedFootForm"`
	StuffingPaper             string `gorm:"column:StuffingPaper"`
}
