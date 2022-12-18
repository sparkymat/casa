package main

import (
	"git.orion.home/oxhead/casa/config"
	"git.orion.home/oxhead/casa/database"
	"git.orion.home/oxhead/casa/route"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := database.New(database.Config{DataPath: cfg.DataPath()})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	route.Setup(e, cfg, db)
	e.Logger.Fatal(e.Start(":8080"))
}
