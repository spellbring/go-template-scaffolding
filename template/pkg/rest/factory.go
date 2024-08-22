package rest

import (
	"{{bootstrap_template}}/pkg/log/logger"
	"errors"
	"time"
)

type Client interface {
	Post(string, interface{}, map[string]string) (int, []byte, error)
	Get(string, map[string]string) (int, []byte, error)
	GetWithParams(string, map[string]string, map[string]string) (int, []byte, error)
}

var (
	errInvalidRestClientInstance = errors.New("invalid client instance")
)

const (
	InstanceRestyV2 int = iota
)

func NewRestClientFactory(
	instance int,
	log logger.Logger,
	ctxTimeout time.Duration,
) (Client, error) {
	switch instance {
	case InstanceRestyV2:
		return newRestyV2(restClientConfig(), log), nil
	default:
		return nil, errInvalidRestClientInstance
	}
}
