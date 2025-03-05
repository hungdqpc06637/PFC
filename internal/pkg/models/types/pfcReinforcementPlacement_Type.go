package types

type PFC_ReinforcementPlacement struct {
	ReinforcementPlacementID     string                           `gorm:"column:ReinforcementPlacementID"`
	ModelType                    string                           `gorm:"column:ModelType"`
	ModelName                    string                           `gorm:"column:ModelName"`
	MaterialNumber               string                           `gorm:"column:MaterialNumber"`
	Title                        string                           `gorm:"column:Title"`
	ItemIndex                    string                           `gorm:"column:ItemIndex"`
}

type PFC_ItemReinforcementPlacement struct {
	ItemReinforcementPlacementID string `gorm:"column:ItemReinforcementPlacementID"`
	ReinforcementPlacementID     string `gorm:"column:ReinforcementPlacementID"`
	Component                    string `gorm:"column:Component"`
	ImageContent                 string `gorm:"column:ImageContent"`
	Material                     string `gorm:"column:Material"`
	Adhesive                     string `gorm:"column:Adhesive"`
	AttachingMethod              string `gorm:"column:AttachingMethod"`
	Temp                         string `gorm:"column:Temp"`
	Pressure                     string `gorm:"column:Pressure"`
	Time                         string `gorm:"column:Time"`
	ItemIndex                    string `gorm:"column:ItemIndex"`
}
