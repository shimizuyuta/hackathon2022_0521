package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/squadra-ricordo/ent"
	"github.com/squadra-ricordo/ent/migrate"
)

func main() {
	client, err := ent.Open("mysql", "root:Pvmgz-mysql-1276@tcp(127.0.0.1:3306)/test_go")
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
