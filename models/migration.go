package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sintell/mmo-server/utils"
	"os"
)

var db gorm.DB
var dbConnString string

func init() {
	settings := utils.Settings{}
	settings.LoadArgs()

	dbConnString = fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s",
		settings.DbUser, settings.DbName, settings.DbPass)

	fmt.Println(dbConnString)

	db = GetDB()
	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Character{})

	db.AutoMigrate(&User{}, &Character{})
	db.Model(&User{}).AddIndex("idx_user_name", "name", "deleted_at")
	db.Model(&User{}).AddIndex("idx_user_uid", "uid", "deleted_at")

	db.Model(&Character{}).AddIndex("idx_character_name", "name", "deleted_at")
	db.Model(&Character{}).AddIndex("idx_character_class", "class", "deleted_at")
	db.Model(&Character{}).AddIndex("idx_character_level", "level", "deleted_at")
}

func GetDB() gorm.DB {
	DB, err := gorm.Open("postgres", dbConnString)

	if err != nil {
		panic(err.Error())
		os.Exit(2)
	}

	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(50)
	DB.DB().SetMaxOpenConns(300)

	return DB
}
