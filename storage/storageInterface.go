package storage

import (
	"github.com/EmiSan1998/gpsTiming-backend/datatypes"
	"github.com/google/uuid"
)

// Store Interface to interact with storage of routes
type Store interface {
	GetRoute(id string) (route datatypes.Route, exists bool)
	CreateRoute(route datatypes.Route) uuid.UUID
}
