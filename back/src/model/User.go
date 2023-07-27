package model

type User struct {
	ID       uint `gorm:"primarykey"`
	Password string
	Email    string
	captcha  string
}

func (*User) TableName() string {
	return "USER"
}
