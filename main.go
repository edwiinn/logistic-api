package main

import (
	"logistic-api/core"

	_ "logistic-api/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	core.Run()
}
