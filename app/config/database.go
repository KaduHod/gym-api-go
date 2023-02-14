package config

import (
	"api/app/helpers/errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/gymapp?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errors.CheckPanic(err)
	return Db
}
