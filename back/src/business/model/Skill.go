package model

type Skill struct {
	Type string `json:"type,omitempty"`
	Img  string `json:"img,omitempty"`
}

func MakeSkill(_type string, img string) Skill {
	return Skill{Type: _type, Img: img}
}
