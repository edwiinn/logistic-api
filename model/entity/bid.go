package entity

type Bid struct {
	ID                int
	JobID             int `sql:"type:int REFERENCES job(id)"`
	TransporterName   string
	TransporterRating string
	Price             float64
	VehicleName       string
}
