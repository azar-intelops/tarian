package store

import "github.com/devopstoday11/tarian/pkg/tarianpb"

type EventStore interface {
	GetAll() ([]*tarianpb.Event, error)
	FindByNamespace(namespace string) ([]*tarianpb.Event, error)
	Add(*tarianpb.Event) error
}