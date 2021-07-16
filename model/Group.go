package model

import (
	"log"
	"time"
	"xietong.me/ginessential/common"
)

type Group struct {
	GroupId   uint      `gorm:"primaryKey;" json:"groupId" uri:"groupId"`
	GroupName string    `gorm:"varchar(255);unique" json:"groupName" uri:"groupName"`
	ItemCOUNT int       `json:"item" uri:"item"`
	User      User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId    int       `json:"userId" uri:"userId"`
	CreatedAt time.Time `json:"createdAt" uri:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" uri:"updatedAt" gorm:"autoUpdateTime"`
}

func (*Group) FindALlGroup() []Group {
	db := common.GetDB()
	var groupList []Group
	err := db.Find(&Group{}, &groupList)
	if err != nil {
		return []Group{}
	} else {
		return groupList
	}
}
func (*Group) FindGroupByUserName(userName string) []Group {
	db := common.GetDB()
	var user User
	err := db.Where(&User{Name: userName}).First(&user).Error
	if err == nil {
		return []Group{}
	} else {
		var groups []Group
		err := db.Where("userId = ?", user.UserId).Find(&groups)
		if err == nil {
			return []Group{}
		} else {
			return groups
		}
	}
}

func (*Group) UpdateGroup(g Group) error {
	db := common.GetDB()
	err := db.Model(&Group{}).Where("groupId = ?", g.GroupId).Update("groupName = ?", g.GroupName).Error
	if err != nil {
		log.Println(err.Error())
		return err
	} else {
		return err
	}
}

func (*Group) DeleteGroup(g Group) error {
	groupId := g.GroupId
	db := common.GetDB()
	err := db.Model(&Group{}).Where("groupId = ?", groupId).Delete(Group{}).Error
	if err != nil {
		log.Println(err.Error())
		return err
	} else {
		return err
	}
}
