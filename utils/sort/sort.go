package sort

import (
	"logistic-api/model/entity"
)

func Sort(arr SortPropertiesInterface) {
	quickSort(arr, 0, arr.Length()-1)
}

func JobByOrigin(arr []entity.Job, order OrderBy) {
	if order == ASCENDING {
		Sort(jobByOriginProperties(arr))
	} else {
		Sort(Reverse(jobByOriginProperties(arr)))
	}
}
func JobByDestination(arr []entity.Job, order OrderBy) {
	if order == ASCENDING {
		Sort(jobByDestinationProperties(arr))
	} else {
		Sort(Reverse(jobByDestinationProperties(arr)))
	}
}
func JobByBudget(arr []entity.Job, order OrderBy) {
	if order == ASCENDING {
		Sort(jobByBudgetProperties(arr))
	} else {
		Sort(Reverse(jobByBudgetProperties(arr)))
	}
}
func JobByShipment(arr []entity.Job, order OrderBy) {
	if order == ASCENDING {
		Sort(jobByShipmentProperties(arr))
	} else {
		Sort(Reverse(jobByShipmentProperties(arr)))
	}
}
func JobByDistance(arr []entity.Job, order OrderBy) {
	if order == ASCENDING {
		Sort(jobByDistanceProperties(arr))
	} else {
		Sort(Reverse(jobByDistanceProperties(arr)))
	}
}
func BidByTransporterName(arr []entity.Bid, order OrderBy) {
	if order == ASCENDING {
		Sort(bidByTransporterNameProperties(arr))
	} else {
		Sort(Reverse(bidByTransporterNameProperties(arr)))
	}
}
func BidByTransporterRating(arr []entity.Bid, order OrderBy) {
	if order == ASCENDING {
		Sort(bidByTransporterRatingProperties(arr))
	} else {
		Sort(Reverse(bidByTransporterRatingProperties(arr)))
	}
}
func BidByPrice(arr []entity.Bid, order OrderBy) {
	if order == ASCENDING {
		Sort(bidByPriceProperties(arr))
	} else {
		Sort(Reverse(bidByPriceProperties(arr)))
	}
}
func BidByVehicleName(arr []entity.Bid, order OrderBy) {
	if order == ASCENDING {
		Sort(bidByVehicleNameProperties(arr))
	} else {
		Sort(Reverse(bidByVehicleNameProperties(arr)))
	}
}

//https://www.geeksforgeeks.org/quick-sort/
func quickSort(arr SortPropertiesInterface, low, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr SortPropertiesInterface, low, high int) int {
	pivot := high
	i := low - 1
	for j := low; j < high; j++ {
		if arr.Less(j, pivot) {
			i++
			arr.Swap(i, j)
		}
	}
	arr.Swap(i+1, high)
	return i + 1
}
