package entity

import (
	"time"
)

type Job struct {
	Id          int
	Origin      string
	Destination string
	Budget      float64
	Shipment    time.Time
	Distance    float64
}
