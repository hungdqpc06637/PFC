package types

type PFC_AdhesiveOtherType struct {
	AdhesiveOtherTypeID string `gorm:"column:AdhesiveOtherTypeID"`
	LaminationProcessID string `gorm:"column:LaminationProcessID"`
	Name                string `gorm:"column:Name"`
	Description         string `gorm:"column:Description"`
}
