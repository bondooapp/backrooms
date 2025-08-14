package database

import (
	"context"
	"fmt"
	"github.com/bondooapp/backrooms/util"
	"github.com/bondooapp/backrooms/util/xlog"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresParam struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type PostgresClient struct {
	Client *gorm.DB
}

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

func NewPostgresClient(cfg *PostgresParam) (*PostgresClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		xlog.Fatal(context.Background(), err, "connection to postgres failed")
		return nil, err
	}
	return &PostgresClient{Client: db}, nil
}
