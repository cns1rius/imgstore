package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

type User struct {
	ID       uint   `gorm:"primarykey"`
	UserName string `gorm:"unique"`
	Password string
}

type Img struct {
	ImgID      uint   `gorm:"primarykey"`
	Path       string `gorm:"unique"`
	Permission string `gorm:"not null"`
}

func InitDB() {
	//driverName := Conf.GetString("db.DriverName")
	host := Conf.GetString("db.Host")
	port := Conf.GetString("db.Port")
	username := Conf.GetString("db.Username")
	password := Conf.GetString("db.Password")
	database := Conf.GetString("db.Database")
	charset := Conf.GetString("db.Charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gin_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	if err = db.AutoMigrate(&User{}, &Img{}); err != nil {
		panic(err)
	}
	DB = db
}
