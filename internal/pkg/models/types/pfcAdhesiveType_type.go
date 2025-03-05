package types

type PFC_AdhesiveType struct {
	AdhesiveTypeID      string `gorm:"column:AdhesiveTypeID"`
	LaminationProcessID string `gorm:"column:LaminationProcessID"`
	Type                string `gorm:"column:Type"`
	Name                string `gorm:"column:Name"`
	Vendor              string `gorm:"column:Vendor"`
	Thickness           string `gorm:"column:Thickness"`
	MeltingPoint        string `gorm:"column:MeltingPoint"`
}
