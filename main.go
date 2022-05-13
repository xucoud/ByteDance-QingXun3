package main

import (
	"context"
	"fmt"
	"gorm-gen-peoject/config"
	"gorm-gen-peoject/db"
)

func main() {
	// 初始化配置
	config.InitConfig()
	//初始化数据库连接
	db.InitDb()
	// 结构体转数据库表
	struct2db()
	// 数据库表生成CRUD
	table2struct()

	//添加100个用户信息
	AddPeople(context.Background())

	//获取版本最大的用户信息
	fmt.Println("版本数最大的用户数量为：",GetMaxVersionCount(context.Background()))
}
