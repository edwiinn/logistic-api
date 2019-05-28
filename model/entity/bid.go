package entity

type Bid struct {
	ID                int     `json:"id"`
	JobID             int     `json:"job_id"`
	TransporterName   string  `json:"transporter_name"`
	TransporterRating string  `json:"transporter_rating"`
	Price             float64 `json:"price"`
	VehicleName       string  `json:"vehicle_name"`
}
