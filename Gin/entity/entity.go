package entity

//user entity
type User struct {
	UserName string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Enabled  bool   `gorm:"column:enabled"`
}

//
func (User) TableName() string {
	return "users"
}
