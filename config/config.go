package config

import (
	"os"

	"github.com/joho/godotenv"
)

var DatabaseDriver string
var DatabaseSource string

func init() {
	godotenv.Load(".env")
	DatabaseDriver = os.Getenv("DATABASE_DRIVER")
	databaseUserName := os.Getenv("DATABASE_USER")
	databaseUserPass := os.Getenv("DATABASE_PASS")
	databaseName := os.Getenv("DATABASE_NAME")
	DatabaseSource = databaseUserName + ":" + databaseUserPass + "@/" + databaseName + "?charset=utf8&parseTime=True&loc=Local"
}
