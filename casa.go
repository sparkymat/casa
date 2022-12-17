package main

import (
	"git.orion.home/oxhead/casa/config"
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

	e := echo.New()
	route.Setup(e, cfg)
	e.Logger.Fatal(e.Start(":8080"))
}
