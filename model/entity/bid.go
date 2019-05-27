package entity

type Bid struct {
	ID                int
	JobID             int
	TransporterName   string
	TransporterRating string
	Price             float64
	VehicleName       string
}
