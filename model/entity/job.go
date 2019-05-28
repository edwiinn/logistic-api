package entity

import (
	"time"
)

type Job struct {
	ID          int       `json:"id"`
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Budget      float64   `json:"budget"`
	Shipment    time.Time `json:"shipment"`
	Distance    float64   `json:"distance"`
}
