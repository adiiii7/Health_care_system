package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "username:password@tcp(host:port)/dbname?parseTime=true"
	// Replace 'username', 'password', 'host', 'port', and 'dbname' with your MySQL database credentials

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
}

func CloseDB() {
	if DB != nil {
		db, _ := DB.DB()
		db.Close()
	}
}
