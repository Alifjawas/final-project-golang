package database

import (
	"final-project-gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

//func init
func StartDB() {
	//todo
	// dsn
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "alif:alif@tcp(127.0.0.1:3306)/final_project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//open database
	db = db.Debug()
	autoMigrate()
}

func GetDB() *gorm.DB {
	return db
}

func autoMigrate() {
	db.AutoMigrate(models.User{
		GormModel: models.GormModel{},
		FullName:  "",
		Email:     "",
		Password:  "",
	})
}
