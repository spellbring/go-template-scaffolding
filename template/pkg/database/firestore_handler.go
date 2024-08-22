package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

type firestoreHandler struct {
	client *firestore.Client
}

func NewFirestoreHandler(c *config) (*firestoreHandler, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, c.projectId)
	if err != nil {
		log.Fatal(err)
	}

	return &firestoreHandler{
		client: client,
	}, nil
}
