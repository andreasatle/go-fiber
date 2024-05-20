package database

import (
	"fmt"
	"log"

	"github.com/andreasatle/go-fiber/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email string
}

var DB *gorm.DB

// ConnectDB connects to the database
func init() {
	var err error
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Host, config.Cfg.Database.Port, config.Cfg.Database.Name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	DB.AutoMigrate(&User{})
	//password1, err := crypt.HashPassword("password1")
	//password2, err := crypt.HashPassword("password2")
	//password3, err := crypt.HashPassword("password3")
	//password4, err := crypt.HashPassword("password4")
	//user1 := User{Username: "user1", Password: string(password1), Email: "email1@example.com"}
	//user2 := User{Username: "user2", Password: string(password2), Email: "email2@example.com"}
	//user3 := User{Username: "user3", Password: string(password3), Email: "email3@example.com"}
	//user4 := User{Username: "user4", Password: string(password4), Email: "email4@example.com"}
	//DB.Create(&user1)
	//DB.Create(&user2)
	//DB.Create(&user3)
	//DB.Create(&user4)
}