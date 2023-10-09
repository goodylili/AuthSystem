package database

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

type Database struct {
	Client *gorm.DB
	Logger *logrus.Logger
}

func NewDatabase() (*Database, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	envVars := []string{"DB_HOST", "DB_PORT", "DB_USERNAME", "DB_TABLE", "DB_PASSWORD", "SSL_MODE"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			return nil, fmt.Errorf("environment variable %s not set", envVar)
		}
	}

	configurations := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_TABLE"), os.Getenv("DB_PASSWORD"), os.Getenv("SSL_MODE"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  configurations,
		PreferSimpleProtocol: true,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Enable connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying SQL DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return &Database{
		Client: db,
		Logger: logger,
	}, nil
}

// Ping - pings the database to check if it is alive
func (d *Database) Ping(ctx context.Context) error {
	client, err := d.Client.DB()
	if err != nil {
		return err
	}
	return client.PingContext(ctx)
}
