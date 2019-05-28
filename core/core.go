package core

import (
	"logistic-api/config"
	"logistic-api/model/entity"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var engine *echo.Echo = echo.New()

func Run() {
	setupDatabase()
	engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	defer engine.Close()
	engine.Logger.Fatal(engine.Start(config.ServerURL))
}

func setupDatabase() {
	db, err := gorm.Open(config.DatabaseDriver, config.DatabaseSource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Job{})
	db.AutoMigrate(&entity.Bid{}).AddForeignKey("job_id", "jobs(id)", "NO ACTION", "NO ACTION")
}
