package database

import (
	"errors"
)

var (
	errInvalidSQLDatabaseInstance = errors.New("invalid connsql db instance")
)

const (
	InstancePostgres int = iota
)

func NewDatabaseSQLFactory(instance int) (IPostgresHandler, error) {
	switch instance {
	case InstancePostgres:
		return NewPostgresHandler(newConfigPostgres())
	default:
		return nil, errInvalidSQLDatabaseInstance
	}
}
