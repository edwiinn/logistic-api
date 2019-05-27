package sort

import (
	"logistic-api/model/entity"
)

type SortPropertiesInterface interface {
	Length() int
	Less(i, j int) bool
	Swap(i, j int)
}

type jobByOriginProperties []entity.Job
type jobByDestinationProperties []entity.Job
type jobByBudgetProperties []entity.Job
type jobByShipmentProperties []entity.Job
type jobByDistanceProperties []entity.Job
type bidByTransporterNameProperties []entity.Bid
type bidByTransporterRatingProperties []entity.Bid
type bidByPriceProperties []entity.Bid
type bidByVehicleNameProperties []entity.Bid

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
func (jb jobByOriginProperties) Length() int {
	return len(jb)
}

func (jb jobByOriginProperties) Less(i, j int) bool {
	if jb[i].Origin <= jb[j].Origin {
		return true
	}
	return false
}

func (jb jobByOriginProperties) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

//Job By Destination
func (jb jobByDestinationProperties) Length() int {
	return len(jb)
}

func (jb jobByDestinationProperties) Less(i, j int) bool {
	if jb[i].Destination <= jb[j].Destination {
		return true
	}
	return false
}
func (jb jobByDestinationProperties) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

//Job By Budget
func (jb jobByBudgetProperties) Length() int {
	return len(jb)
}
func (jb jobByBudgetProperties) Less(i, j int) bool {
	if jb[i].Budget <= jb[j].Budget {
		return true
	}
	return false
}
func (jb jobByBudgetProperties) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

//Job By Shipment
func (jb jobByShipmentProperties) Length() int {
	return len(jb)
}
func (jb jobByShipmentProperties) Less(i, j int) bool {
	return jb[i].Shipment.Before(jb[j].Shipment)
}
func (jb jobByShipmentProperties) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

//Job By Distance
func (jb jobByDistanceProperties) Length() int {
	return len(jb)
}
func (jb jobByDistanceProperties) Less(i, j int) bool {
	if jb[i].Distance <= jb[j].Distance {
		return true
	}
	return false
}
func (jb jobByDistanceProperties) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

//Bid By Transporter Name
func (bb bidByTransporterNameProperties) Length() int {
	return len(bb)
}
func (bb bidByTransporterNameProperties) Less(i, j int) bool {
	if bb[i].TransporterName <= bb[j].TransporterName {
		return true
	}
	return false
}
func (bb bidByTransporterNameProperties) Swap(i, j int) {
	bb[i], bb[j] = bb[j], bb[i]
}

//Bid By Transporter Rating
func (bb bidByTransporterRatingProperties) Length() int {
	return len(bb)
}
func (bb bidByTransporterRatingProperties) Less(i, j int) bool {
	if bb[i].TransporterRating <= bb[j].TransporterRating {
		return true
	}
	return false
}
func (bb bidByTransporterRatingProperties) Swap(i, j int) {
	bb[i], bb[j] = bb[j], bb[i]
}

//Bid By Price
func (bb bidByPriceProperties) Length() int {
	return len(bb)
}
func (bb bidByPriceProperties) Less(i, j int) bool {
	if bb[i].Price <= bb[j].Price {
		return true
	}
	return false
}
func (bb bidByPriceProperties) Swap(i, j int) {
	bb[i], bb[j] = bb[j], bb[i]
}

//Bid by Vehicle Name
func (bb bidByVehicleNameProperties) Length() int {
	return len(bb)
}
func (bb bidByVehicleNameProperties) Less(i, j int) bool {
	if bb[i].VehicleName <= bb[j].VehicleName {
		return true
	}
	return false
}
func (bb bidByVehicleNameProperties) Swap(i, j int) {
	bb[i], bb[j] = bb[j], bb[i]
}
