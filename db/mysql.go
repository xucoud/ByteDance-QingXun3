package db

import (
	"fmt"
	"gorm-gen-peoject/config"
	"gorm-gen-peoject/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDb() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		config.Config.UserName,
		config.Config.Password,
		config.Config.Host,
		config.Config.Port,
		config.Config.DBName,
		config.Config.Charset)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println("open MYSQL fail, err =", err.Error())
		return
	}
	db.Where("version < 100").Delete(&model.User{})
	DB = db
}
