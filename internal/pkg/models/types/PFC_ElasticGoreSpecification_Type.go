package types

type PFC_ElasticGoreSpecification struct {
	ElasticGoreSpecificationID string `gorm:"column:ElasticGoreSpecificationID"`
	ModelType                  string `gorm:"column:ModelType"`
	ModelName                  string `gorm:"column:ModelName"`
	MaterialNumber             string `gorm:"column:MaterialNumber"`
	Title                      string `gorm:"column:Title"`
	ItemIndex                  string `gorm:"column:ItemIndex"`
}

type PFC_ItemElasticGoreSpecification struct {
	ItemElasticGoreSpecificationID string `gorm:"column:ItemElasticGoreSpecificationID"`
	ElasticGoreSpecificationID     string `gorm:"column:ElasticGoreSpecificationID"`
	Component                      string `gorm:"column:Component"`
	ImageContent                   string `gorm:"column:ImageContent"`
	Vendor                         string `gorm:"column:Vendor"`
	Material                       string `gorm:"column:Material"`
	Model                          string `gorm:"column:Model"`
	TableRow1                      string `gorm:"column:TableRow1"`
	TableRow2                      string `gorm:"column:TableRow2"`
	TableRow3                      string `gorm:"column:TableRow3"`
	TableRow4                      string `gorm:"column:TableRow4"`
	TableRow5                      string `gorm:"column:TableRow5"`
	ItemIndex                      string `gorm:"column:ItemIndex"`
}
