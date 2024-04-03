package cmd

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
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
		sslMode := os.Getenv("DB_SSL_MODE")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s sslmode=%s", host, port, user, password, dbname, timezone, sslMode)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			slog.Error("Failed to connect to database", "error", err)
			return
		}

		c.db = db
		slog.Info("Connected to database", "dbname", dbname)
		if err != nil {
			c.err = err
			return
		}
	})

	return c.db, c.err
}

func GetDB() *gorm.DB {
	db, err := dbInstance.getDB()
	if err != nil {
		slog.Error("Failed to connect to database: %s", err)
		return nil
	}
	return db
}
