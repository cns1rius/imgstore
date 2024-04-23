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
	gorm.Model
	// ID       uint   `gorm:"primarykey"`
	UserName string `gorm:"unique"`
	Password string
}

type Img struct {
	gorm.Model
	// ImgID      uint   `gorm:"primarykey"`
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

func ImgPermissionImg(perm string) []Img {
	var (
		ImgTable []Img
	)
	if strings.Contains(Conf.GetString("admin"), perm) {
		_ = DB.Raw("SELECT * FROM gin_img").Scan(&ImgTable).Error
		return ImgTable
	} else {
		_ = DB.Raw("SELECT * FROM gin_img WHERE FIND_IN_SET(?, permission)", perm).Scan(&ImgTable).Error
		return ImgTable
	}
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
	} else if !containsPerm(imgPathPermission(path), perm) {
		perm = imgPathPermission(path) + perm + ","
		err := DB.Model(&ImgTable).Updates(Img{Path: path, Permission: perm}).Error
		return err
	}
	return errors.New("already in your store")
}

func ImgDelete(imgIDList []int, userID string) {
	var ImgTable Img

	if strings.Contains(Conf.GetString("admin"), userID) {
		DB.Delete(&ImgTable, imgIDList)
	} else {
		for _, id := range imgIDList {
			_ = DB.Where("id = ?", id).First(&ImgTable)
			newPermission := removeIDFromPermission(ImgTable.Permission, userID)
			if newPermission != "" {
				_ = DB.Model(&ImgTable).Updates(Img{Path: ImgTable.Path, Permission: newPermission}).Error
			} else {
				DB.Delete(&ImgTable, id)
			}
		}
	}
}

// 辅助函数：从逗号分隔的字符串中移除指定的ID
func removeIDFromPermission(permission string, perm string) string {
	ids := strings.Split(permission, ",")
	result := make([]string, 0, len(ids))

	for _, i := range ids {
		if perm != i {
			result = append(result, i)
		}
	}

	return strings.Join(result, ",")
}

// 辅助函数：判断perm是否在permission中
func containsPerm(permission string, perm string) bool {
	ids := strings.Split(permission, ",")
	for _, i := range ids {
		if perm == i {
			return true
		}
	}
	return false
}
