package publisher

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

type PublisherClient struct {
	topic *pubsub.Topic
	ctx   context.Context
}

type IPublisherClient interface {
	PublishEvent(message []byte, att map[string]string) (string, error)
}

func NewPublisherClient() (*PublisherClient, error) {
	ctx := context.Background()
	conf := GenericPublisherConfig()
	client, err := pubsub.NewClient(ctx, conf.ProjectID)

	if err != nil {
		return nil, err
	}

	topic := client.Topic(conf.TopicName)
	if topic == nil {
		return nil, errors.New("topic does not exist: " + conf.TopicName)
	}

	topic.EnableMessageOrdering = true

	return &PublisherClient{
		topic: topic,
		ctx:   ctx,
	}, nil
}

func (d *PublisherClient) PublishEvent(message []byte, att map[string]string) (string, error) {

	now := time.Now()
	att["timestamp"] = fmt.Sprintf("%d", now.UnixNano())
	att["datetime"] = now.UTC().Format(time.RFC3339)

	result := d.topic.Publish(d.ctx, &pubsub.Message{
		Data:        message,
		OrderingKey: att["timestamp"],
		Attributes:  att,
	})

	id, err := result.Get(d.ctx)

	if err != nil {
		return "", errors.New("an error has occurred publishing the message" + err.Error())
	}

	return id, nil
}
