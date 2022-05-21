package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/squadra-ricordo/ent"
	lib "github.com/squadra-ricordo/lib"
)

var client *ent.Client
var ctx context.Context

func main() {
	lib.InitDB()

	mysqlConfig := mysql.Config{
		User:                 "root",
		Passwd:               os.Getenv("PASSWORD"),
		Addr:                 os.Getenv("DB_HOST") + ":3306",
		DBName:               os.Getenv("DB_NAME"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	}
	client, err := ent.Open("mysql", mysqlConfig.FormatDSN())

	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password_hash := c.FormValue("password_hash")

	lib.CreateUser(client, name, email, password_hash, ctx)

	return c.String(http.StatusOK, "")
}
