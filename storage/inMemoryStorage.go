package storage

import (
	"github.com/EmiSan1998/gpsTiming-backend/datatypes"
	"github.com/google/uuid"
)

// InMemoryStorage simple in-memory storage, to be used only for debugging/testing purposes
type InMemoryStorage struct {
	Routes map[string]datatypes.Route
}

//GetRoute ...
func (storage InMemoryStorage) GetRoute(id string) (route datatypes.Route, exists bool) {
	route, exists = storage.Routes[id]
	return route, exists
}

//CreateRoute ...
func (storage InMemoryStorage) CreateRoute(route datatypes.Route) uuid.UUID {
	randomKey, _ := uuid.NewRandom()
	storage.Routes[randomKey.String()] = route

	return randomKey
}
