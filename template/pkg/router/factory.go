package router

import (
	"errors"
	"time"
	"net/http"

	"{{bootstrap_template}}/pkg/log/logger"
)

type Server interface {
	GetHttpServer() *http.Server
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceGorillaMux int = iota
)

func NewWebServerFactory(instance int, log logger.Logger, port Port, ctxTimeout time.Duration) (Server, error) {
	switch instance {
	case InstanceGorillaMux:
		return newGorillaMux(log, port, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
