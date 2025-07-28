package main

import (
	"fmt"

	"gorm_learn/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	username := "root"        //账号
	password := "123456"      //密码
	host := "192.168.200.128" //数据库地址(本地测试用)
	port := 3306              //数据库端口
	Dbname := "gorm_learn"    //数据库名
	timeout := "10s"          //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	var err error
	db, err = gorm.Open(mysql.Open(dsn)) //不要使用:=，否则会导致db为局部变量
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	fmt.Println(db)

	db.AutoMigrate(&models.User{}) //见models的定义

}
