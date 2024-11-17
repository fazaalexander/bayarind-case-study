package mysql

import (
	"fmt"

	"github.com/fazaalexander/bayarind-case-study.git/config"
	"github.com/fazaalexander/bayarind-case-study.git/database/seed"
	ae "github.com/fazaalexander/bayarind-case-study.git/modules/entity/author"
	be "github.com/fazaalexander/bayarind-case-study.git/modules/entity/book"
	ue "github.com/fazaalexander/bayarind-case-study.git/modules/entity/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
	seed.DBSeed(DB)
}

func InitDB() {
	var err error

	configurations := config.GetConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configurations.DB_USERNAME,
		configurations.DB_PASSWORD,
		configurations.DB_HOST,
		configurations.DB_PORT,
		configurations.DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		ae.Author{},
		be.Book{},
		ue.User{},
	)
}
