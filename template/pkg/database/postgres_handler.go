package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IPostgresHandler interface {
	GetConnection() *gorm.DB
	Migrate()
}

type PostgresHandler struct {
	db *gorm.DB
}

func NewPostgresHandler(c *config) (*PostgresHandler, error) {
	var ds = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.host,
		c.port,
		c.user,
		c.database,
		c.password,
	)
	db, err := gorm.Open(postgres.Open(ds), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(5)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(20 * time.Minute)

	return &PostgresHandler{db}, nil
}

func (postgresHandler PostgresHandler) GetConnection() *gorm.DB {
	return postgresHandler.db
}

func (postgresHandler PostgresHandler) Migrate() {

}
