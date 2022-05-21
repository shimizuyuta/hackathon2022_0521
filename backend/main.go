package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/squadra-ricordo/ent"
	lib "github.com/squadra-ricordo/lib"
)

type UserPostReqBody struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func main() {
	fmt.Println("Please wait...")

	time.Sleep(time.Second * 10)

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

	e.POST("/users", func(c echo.Context) error {
		p := &UserPostReqBody{}
		err := c.Bind(p)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		ctx := context.Background()

		lib.CreateUser(client, p.Name, p.Email, p.PasswordHash, ctx)

		return c.String(http.StatusOK, "")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
