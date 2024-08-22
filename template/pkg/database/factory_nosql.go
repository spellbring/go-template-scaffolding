package database

import (
	"{{bootstrap_template}}/pkg/database/nosql"
	"errors"
)

var (
	errInvalidNoSQLDatabaseInstance = errors.New("invalid nosql db instance")
)

const (
	InstanceFirestoreDB int = iota
)

func NewDatabaseNoSQLFactory(instance int) (nosql.NoSQL, error) {
	switch instance {
	case InstanceFirestoreDB:
		return NewFirestoreHandler(newConfigFirestore())
	default:
		return nil, errInvalidNoSQLDatabaseInstance
	}
}
