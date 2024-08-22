package database

import (
	"os"
	"time"
)

type config struct {
	host       string
	database   string
	port       string
	driver     string
	user       string
	password   string
	collection string
	path       string
	projectId  string

	ctxTimeout time.Duration
}

func newConfigFirestore() *config {
	return &config{
		collection: os.Getenv("FIRESTORE_COLLECTION"),
		path	  : os.Getenv("FIRESTORE_PATH"),
		projectId : os.Getenv("PROJECT_ID"),
		ctxTimeout: 60 * time.Second,
	}
}

func newConfigPostgres() *config {
	return &config{
		host    : os.Getenv("POSTGRES_HOST"),
		database: os.Getenv("POSTGRES_DATABASE"),
		port    : os.Getenv("POSTGRES_PORT"),
		driver  : os.Getenv("POSTGRES_DRIVER"),
		user    : os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
