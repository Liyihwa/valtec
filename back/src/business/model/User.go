package model

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Captcha  string `json:"captcha"`
}

func (*User) TableName() string {
	return "USER"
}
