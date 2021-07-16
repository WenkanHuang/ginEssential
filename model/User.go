package model

import (
	"time"
	"xietong.me/ginessential/common"
)

type User struct {
	UserId    uint      `gorm:"primaryKey;" json:"id" uri:"id"`
	Name      string    `gorm:"varchar(20);not null;unique" json:"name" uri:"name"`
	Password  string    `gorm:"size:255;not null" json:"password" uri:"password"`
	CreatedAt time.Time `json:"createdAt" uri:"createdAt" gorm:"autoCreateTime"`
}

func (*User) SelectAllUsers() []User {
	db := common.GetDB()
	var userList []User
	result := db.Find(&User{}, &userList)
	if result.RowsAffected <= 0 {
		return nil
	} else {
		return userList
	}
}

func (*User) AddUser(user User) error {
	db := common.GetDB()
	result := db.Create(&user).Error
	return result
}

func (*User) FindUserByName(name string) User {
	if name == "" {
		return User{}
	} else {
		db := common.GetDB()
		var user User
		if err := db.Where("name = ?", name).First(&user).Error; err != nil {
			return user
		} else {
			return User{}
		}
	}
}
