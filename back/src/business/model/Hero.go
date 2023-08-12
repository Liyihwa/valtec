package model

type Hero struct {
	ID     uint   `gorm:"primarykey" json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	C      string `json:"c,omitempty"`
	Q      string `json:"q,omitempty"`
	E      string `json:"e,omitempty"`
	X      string `json:"x,omitempty"`
}

func (Hero) TableName() string {
	return "HERO"
}
