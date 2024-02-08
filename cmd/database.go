package cmd

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/lib/pq"
	"ithumans.com/coproxpert/ent"
	"ithumans.com/coproxpert/ent/hook"
	"os"
	"sync"
)

// dbClient implements the Singleton pattern for the database client.
type dbClient struct {
	client *ent.Client
	once   sync.Once
}

var dbInstance = &dbClient{}

func (c *dbClient) getClient(ctx context.Context) (*ent.Client, error) {
	c.once.Do(func() {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		timezone := os.Getenv("DB_TIMEZONE")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s sslmode=disable", host, port, user, password, dbname, timezone)

		var err error
		c.client, err = ent.Open(dialect.Postgres, dsn)
		if err != nil {
			fmt.Printf("Failed to connect to database: %s\n", err)
		}

		// Run the auto migration tool.
		if err := c.client.Schema.Create(ctx); err != nil {
			fmt.Printf("Failed to create schema resources: %s", err)
		}

		// Register hooks
		c.client.Use(hook.UpdateTimestamp)

		fmt.Printf("Connected to database %s\n", dbname)
	})

	return c.client, nil
}

// GetClient returns the singleton instance of the database client.
func GetClient(ctx context.Context) (*ent.Client, error) {
	return dbInstance.getClient(ctx)
}
