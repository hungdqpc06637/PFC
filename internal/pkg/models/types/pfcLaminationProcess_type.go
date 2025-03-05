package types

type PFC_LaminationProcess struct {
	LaminationProcessID  string                    `gorm:"column:LaminationProcessID"`
	ModelType            string                    `gorm:"column:ModelType"`
	ModelName            string                    `gorm:"column:ModelName"`
	MaterialNumber       string                    `gorm:"column:MaterialNumber"`
	ComponentName        string                    `gorm:"column:ComponentName" json:"ComponentName"`
	Vendor               string                    `gorm:"column:Vendor" json:"Vendor"`
	MaterialDescriptions []PFC_MaterialDescription `gorm:"-" json:"MaterialDescriptions"`
	AdhesiveType         []PFC_AdhesiveType        `gorm:"-" json:"AdhesiveType"`
	AdhesiveOtherType    []PFC_AdhesiveOtherType   `gorm:"-" json:"AdhesiveOtherType"`
	Rolls                []PFC_Roll                `gorm:"-" json:"Rolls"`
	Winding              string                    `gorm:"column:Winding" json:"Winding"`
	Cooling              string                    `gorm:"column:Cooling" json:"Cooling"`
	EndStep              string                    `gorm:"column:EndStep" json:"EndStep"`
}
