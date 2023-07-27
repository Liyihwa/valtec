package model

type Map struct {
	ID     uint   `gorm:"primarykey" json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Url    string `json:"url,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

func (*Map) TableName() string {
	return "MAP"
}
