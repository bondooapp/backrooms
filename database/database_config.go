package database

import (
	"context"
	"fmt"
	"time"

	"github.com/bondooapp/backrooms/util"
	"github.com/bondooapp/backrooms/util/xlog"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresParam
//
// Postgres param.
type PostgresParam struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// PostgresClient
//
// Postgres client.
type PostgresClient struct {
	Client *gorm.DB
}

// LoadPostgresParam
//
// Load postgres param.
func LoadPostgresParam() (*PostgresParam, error) {
	_ = godotenv.Load()
	param := &PostgresParam{
		Host:     util.GetEnv("POSTGRES_HOST", "localhost"),
		Port:     util.GetEnv("POSTGRES_PORT", "5432"),
		User:     util.GetEnv("POSTGRES_USER", "root"),
		Password: util.GetEnv("POSTGRES_PASSWORD", "password"),
		DBName:   util.GetEnv("POSTGRES_DB_NAME", "postgres"),
		SSLMode:  util.GetEnv("POSTGRES_SSL_MODE", "disable"),
	}
	return param, nil
}

// NewPostgresClient
//
// New postgres client.
func NewPostgresClient(ctx context.Context, pp *PostgresParam) (*PostgresClient, error) {
	// Set postgres connection timeout.
	connCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Get postgres dsn.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		pp.Host,
		pp.User,
		pp.Password,
		pp.DBName,
		pp.Port,
		pp.SSLMode,
	)

	// Configure gorm.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		xlog.Fatal(ctx, err, "backrooms: failed to open postgres dsn")
		return nil, err
	}

	// Configure connection pool by gorm.
	sqlDB, err := db.DB()
	if err != nil {
		xlog.Fatal(ctx, err, "backrooms: failed to get sql DB")
		return nil, err
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Test connection.
	if err := db.WithContext(connCtx).Raw("SELECT 1").Error; err != nil {
		xlog.Error(ctx, err, "backrooms: failed to test postgres connection")
		return nil, err
	}

	return &PostgresClient{Client: db}, nil
}
