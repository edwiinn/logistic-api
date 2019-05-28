package sort

import (
	"logistic-api/model/entity"
)

func Sort(arr SortPropertiesInterface) {
	quickSort(arr, 0, arr.Length()-1)
}

func IsSorted(data SortPropertiesInterface) bool {
	n := data.Length()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

func JobByOrigin(arr []entity.Job, order OrderBy) {
	if order == ASCENDING {
		Sort(JobByOriginProperties(arr))
	} else {
		Sort(Reverse(JobByOriginProperties(arr)))
	}
}
func JobByDestination(arr []entity.Job, order OrderBy) {
	if order == ASCENDING {
		Sort(JobByDestinationProperties(arr))
	} else {
		Sort(Reverse(JobByDestinationProperties(arr)))
	}
}
func JobByBudget(arr []entity.Job, order OrderBy) {
	if order == ASCENDING {
		Sort(JobByBudgetProperties(arr))
	} else {
		Sort(Reverse(JobByBudgetProperties(arr)))
	}
}
func JobByShipment(arr []entity.Job, order OrderBy) {
	if order == ASCENDING {
		Sort(JobByShipmentProperties(arr))
	} else {
		Sort(Reverse(JobByShipmentProperties(arr)))
	}
}
func JobByDistance(arr []entity.Job, order OrderBy) {
	if order == ASCENDING {
		Sort(JobByDistanceProperties(arr))
	} else {
		Sort(Reverse(JobByDistanceProperties(arr)))
	}
}
func BidByTransporterName(arr []entity.Bid, order OrderBy) {
	if order == ASCENDING {
		Sort(BidByTransporterNameProperties(arr))
	} else {
		Sort(Reverse(BidByTransporterNameProperties(arr)))
	}
}
func BidByTransporterRating(arr []entity.Bid, order OrderBy) {
	if order == ASCENDING {
		Sort(BidByTransporterRatingProperties(arr))
	} else {
		Sort(Reverse(BidByTransporterRatingProperties(arr)))
	}
}
func BidByPrice(arr []entity.Bid, order OrderBy) {
	if order == ASCENDING {
		Sort(BidByPriceProperties(arr))
	} else {
		Sort(Reverse(BidByPriceProperties(arr)))
	}
}
func BidByVehicleName(arr []entity.Bid, order OrderBy) {
	if order == ASCENDING {
		Sort(BidByVehicleNameProperties(arr))
	} else {
		Sort(Reverse(BidByVehicleNameProperties(arr)))
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
