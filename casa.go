package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/casa/config"
	"github.com/sparkymat/casa/route"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	route.Setup(e, cfg)
	e.Logger.Fatal(e.Start(":8080"))
}
