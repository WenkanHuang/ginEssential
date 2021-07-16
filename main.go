package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"xietong.me/ginessential/common"
	"xietong.me/ginessential/config"
	_ "xietong.me/ginessential/docs"
)

func main() {
	config.InitConfig()
	db, err := common.InitDB().DB()
	if err != nil {
		log.Print(err.Error())
	} else {
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Print(err.Error())
			}
		}(db)
	}
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
