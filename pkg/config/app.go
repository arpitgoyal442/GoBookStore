package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {

	d, err := gorm.Open("mysql", "root:ArpitGoyal@442@/mytestdatabase?charset=utf8&&parseTime=True&loc=Local")

	

	if err != nil {
		
		fmt.Println("Error in COnnecting to mysql")
		panic(err)
	}

	db = d
}

func GetDb() *gorm.DB {

	return db
}
