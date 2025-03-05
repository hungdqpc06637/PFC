package types

type PFC_MaterialDescription struct {
	MaterialDescriptionID string `gorm:"column:MaterialDescriptionID"`
	LaminationProcessID   string `gorm:"column:LaminationProcessID"`
	Name                  string `gorm:"column:Name"`
	Mat                   string `gorm:"column:Mat"`
}
