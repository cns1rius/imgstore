package config

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
	"strings"
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

// path->path path->permissions permission->path
// 路径获得路径来判断有无 路径获得权限来判断是否需要更新 权限获得路径用于展示图片
func imgPathPath(path string) string {
	var ImgTable Img
	_ = DB.Where("path = ?", path).First(&ImgTable)
	return ImgTable.Path
}

func imgPathPermission(path string) string {
	var ImgTable Img
	_ = DB.Where("path = ?", path).First(&ImgTable)
	return ImgTable.Permission
}

func ImgPermissionPath(perm string) []string {
	var (
		ImgTables []Img
		paths     []string
	)
	DB.Find(&ImgTables)
	for _, i := range ImgTables {
		if strings.Contains(i.Permission, perm) {
			paths = append(paths, i.Path)
		}
	}
	return paths
	//_ = DB.Where("permission = ?", perm).First(&ImgTable)
	//return ImgTable.Permission
}

func ImgUpdate(path string, permission int) error {
	perm := strconv.Itoa(permission)
	ImgTable := Img{
		Path:       path,
		Permission: perm,
	}
	if imgPathPath(path) == "" {
		err := DB.Create(&ImgTable).Error
		return err
	} else if !strings.Contains(imgPathPermission(path), perm) {
		perm = imgPathPermission(path) + perm
		err := DB.Model(&ImgTable).Updates(Img{Path: path, Permission: perm}).Error
		return err
	}
	return errors.New("already in your store")
}
