package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Stu struct{
	gorm.Model
	Name string
	PassWord string
}

var GlobalDb *gorm.DB

func InitModel(){
	db ,err:= gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ihome?parseTime=true")
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}

	//连接池设置
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(30)
	db.DB().SetConnMaxLifetime(60 * 30)

	//设置表名为单数形式
	db.SingularTable(true)
	GlobalDb = db

	//自动迁移  在gorm中建表默认是负数形式
	db.AutoMigrate(new(Stu))

}

func InsertData(){
	var stu Stu
	stu.Name="bj5q"
	stu.PassWord="123456"
	if err :=GlobalDb.Create(&stu);err !=nil  {
		fmt.Println("创建数据失败!")
		return
	}

	fmt.Println(stu)
}
