package lib

import (
	"context"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/squadra-ricordo/ent"
	"github.com/squadra-ricordo/ent/migrate"
)

func InitDB() {
	mysqlConfig := mysql.Config{
		User:                 "root",
		Passwd:               os.Getenv("PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":3306",
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	client, err := ent.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// マイグレーションの実行
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
