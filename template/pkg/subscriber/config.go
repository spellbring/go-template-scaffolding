package subscriber

import "os"

type config struct {
	subscriptionName string
	topicName        string
	projectID        string
}

func genericSubscriptionConfig() *config {
	return &config{
		subscriptionName: os.Getenv("INBOUND_SUBSCRIPTION"),
		projectID:        os.Getenv("PROJECT_ID"),
	}
}
