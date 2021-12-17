package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConnection() *gorm.DB {
	dsn := "root:soscretthatnearlycannotbefindout@tcp(host.docker.internal:3307)/go_image_uploader"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to mysql database")
	}
	return db
}
