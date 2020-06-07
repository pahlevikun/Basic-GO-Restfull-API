package config

import (
	"github.com/jinzhu/gorm"

	"../structs"
)

// DBInit create connection
func DBInit() *gorm.DB {
	db, errors := gorm.Open("mysql", "gouser:123123@tcp(127.0.0.1)/go_test_db")
	if errors != nil {
		panic("Failed connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
