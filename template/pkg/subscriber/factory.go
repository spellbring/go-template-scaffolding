package subscriber

import (
	"{{bootstrap_template}}/pkg/log/logger"
	"context"
	"errors"
)

var (
	errInvalidEventHandlerInstance = errors.New("invalid event handler instance")
)

const (
	InstanceGooglePubSubHandler   int    = iota
	GENERIC_SUBSCRIPTION_INSTANCE string = "GENERIC_SUBSCRIPTION_INSTANCE"
)

type EventHandler interface {
	HandleMessage(ctx context.Context) (string error)
}

func NewEventSubscriberHandlerFactory(
	instance int,
	subType string,
	log logger.Logger,
) (EventHandler, string, error) {
	switch instance {
	case InstanceGooglePubSubHandler:
		if subType == GENERIC_SUBSCRIPTION_INSTANCE {
			return NewPubSubSubscriberHandler(genericSubscriptionConfig(), log, subType)
		} else {
			return nil, "", errors.New("subscription type not found ")
		}
	default:
		return nil, "", errInvalidEventHandlerInstance
	}
}
