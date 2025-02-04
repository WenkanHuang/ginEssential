package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"xietong.me/ginessential/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	fmt.Println(viper.GetString("datasource.driverName"))
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("fail to connect database,err:" + err.Error())
	}
	s := db.AutoMigrate(&model.User{}, &model.Group{}, &model.Todo{})
	if s != nil {
		log.Print(s.Error())
		return nil
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
