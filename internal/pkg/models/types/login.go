package types

// TYPE - Failures
type Login struct {
	UserId   string `gorm:"column:USERID"`
	UserName string `gorm:"column:USERNAME"`
	Password string `gorm:"column:PWD"`
	Role     string
	Token    string
}
