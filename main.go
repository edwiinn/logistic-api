package main

import (
	"logistic-api/config"
	"logistic-api/core"
	"logistic-api/model/entity"

	_ "logistic-api/router"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setupDatabase() {
	db, err := gorm.Open(config.DatabaseDriver, config.DatabaseSource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Bid{}, &entity.Job{})
}

func main() {
	setupDatabase()
	core.Run()
}
