package config

import (
	"github.com/spf13/viper"
	"log"
)

type Mysql struct {
	UserName string `structure:"username"`
	Password string `structure:"password"`
	Host string `structure:"host"`
	Port string `structure:"port"`
	DBName string `structure:"dbname"`
	Charset string `structure:"charset"`
}

var Config Mysql

func InitConfig() {
	// 实例化viper
	v := viper.New()
	//文件的路径如何设置
	v.AddConfigPath("./config")
	v.SetConfigName("setting")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Println("Read Config File Fail，err =", err.Error())
		return
	}
	if err := v.Unmarshal(&Config); err != nil {
		log.Println("Unmarshal Config File Fail，err =", err.Error())
		return
	}
}