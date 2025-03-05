package types

type PFC_UpperMeasurementSpec struct {
	UpperMeasurementSpecID string `gorm:"column:UpperMeasurementSpecID"`
	ModelType              string `gorm:"column:ModelType"`
	ModelName              string `gorm:"column:ModelName"`
	MaterialNumber         string `gorm:"column:MaterialNumber"`
	Title                  string `gorm:"column:Title"`
	ItemIndex              string `gorm:"column:ItemIndex"`
}

type PFC_ItemUpperMeasurementSpec struct {
	ItemUpperMeasurementSpecID string `gorm:"column:ItemUpperMeasurementSpecID"`
	UpperMeasurementSpecID     string `gorm:"column:UpperMeasurementSpecID"`
	ImagesContent              string `gorm:"column:ImagesContent"`
	TableRow1                  string `gorm:"column:TableRow1"`
}
