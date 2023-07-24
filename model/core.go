package model

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
func init()  {
	username:="root"
	password:="cxj2001"
	host:="127.0.0.1"
	port:=3306
	Dbname:="blog_service"

	dsn :=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",username,password,host,port,Dbname) 
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
	})
	if err!=nil {
		log.Fatalln("db connected  error..",err)
	}
    DB=db
  }
