package main

import (
	"context"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/squadra-ricordo/ent"
	"github.com/squadra-ricordo/ent/migrate"
)

func main() {
	client, err := ent.Open("mysql", "root:"+os.Getenv("PASSWORD")+"@tcp("+os.Getenv("LOCALHOST")+":3306)/"+os.Getenv("DB_NAME"))
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
