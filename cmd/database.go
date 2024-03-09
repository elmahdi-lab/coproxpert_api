package cmd

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

// dbClient implements the Singleton pattern for the database client.
type dbClient struct {
	db   *gorm.DB
	once sync.Once
	err  error
}

var dbInstance = &dbClient{}

func (c *dbClient) getDB() (*gorm.DB, error) {
	c.once.Do(func() {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		timezone := os.Getenv("DB_TIMEZONE")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s sslmode=disable", host, port, user, password, dbname, timezone)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("Failed to connect to database: %s\n", err)
			return
		}

		c.db = db
		fmt.Printf("Connected to database %s\n", dbname)

		if err != nil {
			c.err = err
			return
		}
	})

	return c.db, c.err
}

func GetDB() (*gorm.DB, error) {
	db, err := dbInstance.getDB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}
	return db, nil
}
