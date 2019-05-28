package sort

import (
	"logistic-api/model/entity"
)

type SortPropertiesInterface interface {
	Length() int
	Less(i, j int) bool
	Swap(i, j int)
}

type JobByOriginProperties []entity.Job
type JobByDestinationProperties []entity.Job
type JobByBudgetProperties []entity.Job
type JobByShipmentProperties []entity.Job
type JobByDistanceProperties []entity.Job
type BidByTransporterNameProperties []entity.Bid
type BidByTransporterRatingProperties []entity.Bid
type BidByPriceProperties []entity.Bid
type BidByVehicleNameProperties []entity.Bid

//Reverse
type reverse struct {
	SortPropertiesInterface
}

func (r reverse) Less(i, j int) bool {
	return r.SortPropertiesInterface.Less(j, i)
}
func Reverse(data SortPropertiesInterface) SortPropertiesInterface {
	return &reverse{data}
}

//Job By Origin
func (jb JobByOriginProperties) Length() int {
	return len(jb)
}

func (jb JobByOriginProperties) Less(i, j int) bool {
	if jb[i].Origin <= jb[j].Origin {
		return true
	}
	return false
}

func (jb JobByOriginProperties) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

//Job By Destination
func (jb JobByDestinationProperties) Length() int {
	return len(jb)
}

func (jb JobByDestinationProperties) Less(i, j int) bool {
	if jb[i].Destination <= jb[j].Destination {
		return true
	}
	return false
}
func (jb JobByDestinationProperties) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

//Job By Budget
func (jb JobByBudgetProperties) Length() int {
	return len(jb)
}
func (jb JobByBudgetProperties) Less(i, j int) bool {
	if jb[i].Budget <= jb[j].Budget {
		return true
	}
	return false
}
func (jb JobByBudgetProperties) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

//Job By Shipment
func (jb JobByShipmentProperties) Length() int {
	return len(jb)
}
func (jb JobByShipmentProperties) Less(i, j int) bool {
	return jb[i].Shipment.Before(jb[j].Shipment)
}
func (jb JobByShipmentProperties) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

//Job By Distance
func (jb JobByDistanceProperties) Length() int {
	return len(jb)
}
func (jb JobByDistanceProperties) Less(i, j int) bool {
	if jb[i].Distance <= jb[j].Distance {
		return true
	}
	return false
}
func (jb JobByDistanceProperties) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

//Bid By Transporter Name
func (bb BidByTransporterNameProperties) Length() int {
	return len(bb)
}
func (bb BidByTransporterNameProperties) Less(i, j int) bool {
	if bb[i].TransporterName <= bb[j].TransporterName {
		return true
	}
	return false
}
func (bb BidByTransporterNameProperties) Swap(i, j int) {
	bb[i], bb[j] = bb[j], bb[i]
}

//Bid By Transporter Rating
func (bb BidByTransporterRatingProperties) Length() int {
	return len(bb)
}
func (bb BidByTransporterRatingProperties) Less(i, j int) bool {
	if bb[i].TransporterRating <= bb[j].TransporterRating {
		return true
	}
	return false
}
func (bb BidByTransporterRatingProperties) Swap(i, j int) {
	bb[i], bb[j] = bb[j], bb[i]
}

//Bid By Price
func (bb BidByPriceProperties) Length() int {
	return len(bb)
}
func (bb BidByPriceProperties) Less(i, j int) bool {
	if bb[i].Price <= bb[j].Price {
		return true
	}
	return false
}
func (bb BidByPriceProperties) Swap(i, j int) {
	bb[i], bb[j] = bb[j], bb[i]
}

//Bid by Vehicle Name
func (bb BidByVehicleNameProperties) Length() int {
	return len(bb)
}
func (bb BidByVehicleNameProperties) Less(i, j int) bool {
	if bb[i].VehicleName <= bb[j].VehicleName {
		return true
	}
	return false
}
func (bb BidByVehicleNameProperties) Swap(i, j int) {
	bb[i], bb[j] = bb[j], bb[i]
}
