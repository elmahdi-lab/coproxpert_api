package server

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/lib/pq"
	"ithumans.com/coproxpert/ent"
	"ithumans.com/coproxpert/ent/hook"
	"os"
)

var client *ent.Client

func ConnectDatabase() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s sslmode=disable", host, port, user, password, dbname, timezone)

	var err error
	client, err = ent.Open(dialect.Postgres, dsn)
	if err != nil {
		fmt.Printf("Failed to connect to database: %s\n", err)
		return err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Printf("Failed to create schema resources: %s", err)
		return err
	}

	// Register hooks
	client.Use(hook.UpdateTimestamp)

	fmt.Printf("Connected to database %s\n", dbname)

	return nil
}
