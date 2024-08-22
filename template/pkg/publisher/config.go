package publisher

import "os"

type config struct {
	TopicName string
	ProjectID string
}

func GenericPublisherConfig() *config {
	return &config{
		TopicName: os.Getenv("OUTBOUND_TOPIC"),
		ProjectID: os.Getenv("PROJECT_ID"),
	}
}
