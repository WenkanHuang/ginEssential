package model

import (
	"time"
)

type Todo struct {
	TodoID      uint      `gorm:"primaryKey" json:"todoId" uri:"todoId"`
	TodoName    string    `gorm:"primaryKey;unique" json:"todoName" uri:"todoName"`
	TodoContent string    `gorm:"primaryKey" json:"todoContent" uri:"todoContent"`
	IsFinished  bool      `json:"isFinished" uri:"isFinished"`
	UserId      int       `json:"userId" uri:"userId"`
	GroupId     int       `json:"groupId" uri:"groupId"`
	User        User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Group       Group     `gorm:"foreignKey:GroupId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time `json:"createdAt" uri:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" uri:"updatedAt" gorm:"autoUpdateTime"`
}
