package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	lib "github.com/squadra-ricordo/lib"
)

func main() {
	lib.InitDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
