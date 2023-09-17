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

func (h Hero) GetInfo() ResObject {
	resMap := make(ResObject)
	resMap["id"] = h.ID
	resMap["avatar"] = h.Avatar
	resMap["name"] = h.Name
	skills := [4]Skill{MakeSkill("c", h.C), MakeSkill("q", h.Q), MakeSkill("e", h.E), MakeSkill("x", h.X)}
	resMap["skills"] = skills
	return resMap
}
