package main

import (
	"context"
	"fmt"
	"gorm-gen-peoject/db"
	"gorm-gen-peoject/query"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// AddPeople 生成 100个 用户信息
func AddPeople(ctx context.Context) {
	userDao := query.Use(db.DB).User.WithContext(ctx)
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().Unix())
		uuid := rand.Intn(100)
		id := strconv.Itoa(uuid)
		user, err := userDao.FindByUUID(id)
		if err == nil {
			err = userDao.UpdateVersion(id, int(user.Age), user.Name, int(user.Version+1))
			if err != nil {
				//log.Println("update user info fail, err =", err.Error())
				return
			}
		} else {
			err = userDao.InsertValue(id, i, fmt.Sprintf("张三-%v", i), 1)
			if err != nil {
				//log.Println("insert user info fail, err =", err.Error())
				return
			}
		}
	}
}

func GetMaxVersionCount(ctx context.Context) int {
	userDao := query.Use(db.DB).User.WithContext(ctx)
	res, err := userDao.FindMaxVersion()
	if err != nil {
		log.Println("find max version user fail, err =", err.Error())
		return 0
	}
	return res
}
