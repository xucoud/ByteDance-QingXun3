package main

import (
	"fmt"
	"gorm-gen-peoject/db"
	"log"
)

type User struct {
	UUID string
	Name string
	Age int
	Version int
}

func struct2db() {
	err := db.DB.AutoMigrate(&User{})
	if err != nil {
		log.Println("AutoMigrate fail, err =", err.Error())
		return
	}
	fmt.Println("数据迁移成功!")
}
