package subscriber

import (
	"{{bootstrap_template}}/pkg/log/logger"
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
)

type pubsubSubscriberHandler struct {
	client  *pubsub.Client
	sub     *pubsub.Subscription
	log     logger.Logger
	subType string
	sem     chan struct{}
}



func NewPubSubSubscriberHandler(c *config, log logger.Logger, subType string) (*pubsubSubscriberHandler, string, error) {

	ctx := context.Background()

	projectID := c.projectID
	subscriptionID := c.subscriptionName
	client, err := pubsub.NewClient(ctx, projectID)

	if err != nil {
		log.Errorf("Failed to create client: %v", err)
		return nil, "", err
	}

	sub := client.Subscription(subscriptionID)
	exists, err := sub.Exists(ctx)
	if err != nil {
		log.Errorf("Failed to check subscription existence: %v", err)
		return nil, "", err
	}
	if !exists {
		log.Errorf("Subscription %s doesn't exist", subscriptionID)
		return nil, "", errors.New("subscription doesn't exist")
	}

	return &pubsubSubscriberHandler{
		client:  client,
		sub:     sub,
		log:     log,
		subType: subType,
		sem: make(chan struct{}, 1),
	}, c.subscriptionName, nil
}


func (p pubsubSubscriberHandler) HandleMessage(ctx context.Context) error {

	return p.sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		eventSubscriberHandler(msg, &p, ctx)
	})
}

func eventSubscriberHandler(msg *pubsub.Message, p *pubsubSubscriberHandler, ctx context.Context) {

	p.sem <- struct{}{}
	defer func() { <-p.sem }()


	msg.Ack()
}
