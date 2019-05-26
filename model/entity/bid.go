package entity

type Bid struct {
	Id                int
	JobId             int
	TransporterName   string
	TransporterRating string
	Price             float64
}
