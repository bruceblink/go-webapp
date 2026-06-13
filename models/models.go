package models

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Employee struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("custome.db"), &gorm.Config{})
	if err != nil {
		// 打印出绝对具体的错误原因，比如：permission denied
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	err = database.AutoMigrate(&Employee{})
	if err != nil {
		// 不要直接 return，把错误打印出来看看为什么失败
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	DB = database
}
