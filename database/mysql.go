package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dns := "root:@tcp(127.0.0.1:3306)/go_gin_demo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{

	})

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
