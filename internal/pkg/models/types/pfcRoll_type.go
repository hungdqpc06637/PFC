package types

type PFC_Roll struct {
	RollID              string `gorm:"column:RollID"`
	LaminationProcessID string `gorm:"column:LaminationProcessID"`
	Temp                string `gorm:"column:Temp"`
	RollSize            string `gorm:"column:RollSize"`
	Time                string `gorm:"column:Time"`
}
