package dal

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql")

var (
	BookDB *gorm.DB
)

func init(){
	initMysql()
}


func initMysql(){
	var err error
	dsn := "root:zzz991215@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	BookDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	BookDB = BookDB.Table("book")
	fmt.Println(err)
	fmt.Println("sssssssssssssssssss")
	if err != nil{
		panic(err)
	}
}
