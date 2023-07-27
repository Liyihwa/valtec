package model

import "gorm.io/gorm"

type Position struct {
	ID          uint            `form:"id" uri:"id" gorm:"primarykey" json:"id,omitempty"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	HeroID      uint            `form:"hid" uri:"hid" json:"hid,omitempty"`       //所属英雄
	Skill       string          `form:"skill" uri:"skill" json:"skill,omitempty"` //所属技能
	MapID       uint            `form:"mid" uri:"mid" json:"mid,omitempty"`       //所属地图
	StandX      float32         `json:"sx,omitempty"`
	StandY      float32         `json:"sy,omitempty"`
	PutX        float32         `json:"px,omitempty"`
	PutY        float32         `json:"py,omitempty"`
	Like        int             `json:"like,omitempty"`    //点赞数
	Dislike     int             `json:"dislike,omitempty"` //点踩数
	Description string          `json:"description,omitempty"`
}

func (*Position) TableName() string {
	return "POSITION"
}
