package datatypes

import (
	"encoding/json"
	"time"
)

// Coordinate Defines a geographic coordinate
type Coordinate struct {
	Latitude  json.Number `json:"latitude"`
	Longitude json.Number `json:"longitude"`
}

// Segment Defines a segment defined by two couples of coordinates
type Segment struct {
	P1 Coordinate `json:"p1"`
	P2 Coordinate `json:"p2"`
}

// Point Define a point composed by coordinates, altitude and time
type Point struct {
	Position Coordinate  `json:"position"`
	Altitude json.Number `json:"altitude"`
	Time     time.Time   `json:"time"`
}

// Track Defines a track made by a name and an array of points
type Track struct {
	Name string  `json:"name"`
	Path []Point `json:"path"`
}

// Route Defines a route made by a name and an array of segments
type Route struct {
	Name        string    `json:"name"`
	Checkpoints []Segment `json:"checkpoints"`
}
