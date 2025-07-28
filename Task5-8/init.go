/*
连接.go - GORM数据库连接与初始化模块

功能说明:
1. 建立与MySQL数据库的连接
2. 定义User数据模型结构
3. 自动迁移数据库表结构

使用说明:
- 需先配置正确的数据库连接参数
- 依赖gorm.io/gorm和gorm.io/driver/mysql包
*/
package main

import (
	"fmt"

	"gorm_learn/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// init 初始化数据库连接
// 在程序启动时自动执行，建立数据库连接并迁移表结构
func init() {

	username := "root"        //账号
	password := "123456"      //密码
	host := "192.168.200.128" //数据库地址，可以是Ip或者域名
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
	// 连接成功
	fmt.Println(db)

	db.AutoMigrate(&models.User{})

}
