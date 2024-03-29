package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	driverName := Conf.GetString("db.DriverName")
	host := Conf.GetString("db.Host")
	port := Conf.GetString("db.Port")
	username := Conf.GetString("db.Username")
	password := Conf.GetString("db.Password")
	database := Conf.GetString("db.Database")
	charset := Conf.GetString("db.Charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	//db.AutoMigrate(&model.User{})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	return db
}
