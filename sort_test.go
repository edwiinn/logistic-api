package main

import (
	"logistic-api/model/entity"
	"logistic-api/utils/sort"
	"testing"
	"time"
)

func makeDummyJob() []entity.Job {
	jobs := []entity.Job{}
	time1, _ := time.Parse("2006-01-02 15:04:05", "2019-04-18 05:48:18")
	time2, _ := time.Parse("2006-01-02 15:04:05", "2019-04-10 05:48:18")
	time3, _ := time.Parse("2006-01-02 15:04:05", "2019-04-8 05:48:18")
	jobs = append(jobs, entity.Job{
		ID:          1,
		Origin:      "Solo",
		Destination: "Semarang",
		Budget:      5400000,
		Shipment:    time1,
		Distance:    100})
	jobs = append(jobs, entity.Job{
		ID:          2,
		Origin:      "Surabaya",
		Destination: "Jember",
		Budget:      1300000,
		Shipment:    time2,
		Distance:    200})
	jobs = append(jobs, entity.Job{
		ID:          3,
		Origin:      "Bali",
		Destination: "Madura",
		Budget:      1320000,
		Shipment:    time3,
		Distance:    150})
	return jobs
}

func makeDummyBid() []entity.Bid {
	bids := []entity.Bid{}
	bids = append(bids, entity.Bid{
		ID:                1,
		JobID:             1,
		TransporterName:   "Edwin",
		TransporterRating: "Baik",
		Price:             2000000,
		VehicleName:       "Avanza"})
	bids = append(bids, entity.Bid{
		ID:                2,
		JobID:             1,
		TransporterName:   "Hartoyo",
		TransporterRating: "Kurang",
		Price:             2500000,
		VehicleName:       "Toyota"})
	bids = append(bids, entity.Bid{
		ID:                3,
		JobID:             2,
		TransporterName:   "Anthon",
		TransporterRating: "Bagus",
		Price:             2100000,
		VehicleName:       "Suzuki"})
	return bids
}
func TestJobSortByOriginAscending(t *testing.T) {
	jobs := makeDummyJob()
	expectedOutput := []string{"Bali", "Solo", "Surabaya"}
	sort.JobByOrigin(jobs, sort.ASCENDING)
	for k, v := range expectedOutput {
		if jobs[k].Origin != v {
			t.Errorf("Sort By Origin Ascending Failed. Expected : %s ; Actual: %s", v, jobs[k].Origin)
		}
	}
}

func TestJobSortByOriginDescending(t *testing.T) {
	jobs := makeDummyJob()
	expectedOutput := []string{"Surabaya", "Solo", "Bali"}
	sort.JobByOrigin(jobs, sort.DESCENDING)
	for k, v := range expectedOutput {
		if jobs[k].Origin != v {
			t.Errorf("Sort By Origin Descending Failed. Expected : %s ; Actual: %s", v, jobs[k].Origin)
		}
	}
}

func TestJobSortByDestinationAscending(t *testing.T) {
	jobs := makeDummyJob()
	expectedOutput := []string{"Jember", "Madura", "Semarang"}
	sort.JobByDestination(jobs, sort.ASCENDING)
	for k, v := range expectedOutput {
		if jobs[k].Destination != v {
			t.Errorf("Sort By Destination Ascending Failed. Expected : %s ; Actual: %s", v, jobs[k].Destination)
		}
	}
}

func TestJobSortByDestinationDescending(t *testing.T) {
	jobs := makeDummyJob()
	expectedOutput := []string{"Semarang", "Madura", "Jember"}
	sort.JobByDestination(jobs, sort.DESCENDING)
	for k, v := range expectedOutput {
		if jobs[k].Destination != v {
			t.Errorf("Sort By Destination Descending Failed. Expected : %s ; Actual: %s", v, jobs[k].Destination)
		}
	}
}

func TestJobSortByBudgetAscending(t *testing.T) {
	jobs := makeDummyJob()
	expectedOutput := []float64{1300000, 1320000, 5400000}
	sort.JobByBudget(jobs, sort.ASCENDING)
	for k, v := range expectedOutput {
		if jobs[k].Budget != v {
			t.Errorf("Sort By Budget Ascending Failed. Expected : %f ; Actual: %f", v, jobs[k].Budget)
		}
	}
}

func TestJobSortByBudgetDescending(t *testing.T) {
	jobs := makeDummyJob()
	expectedOutput := []float64{5400000, 1320000, 1300000}
	sort.JobByBudget(jobs, sort.DESCENDING)
	for k, v := range expectedOutput {
		if jobs[k].Budget != v {
			t.Errorf("Sort By Budget Descending Failed. Expected : %f ; Actual: %f", v, jobs[k].Budget)
		}
	}
}

func TestJobSortByShipmentAscending(t *testing.T) {
	jobs := makeDummyJob()
	time1, _ := time.Parse("2006-01-02 15:04:05", "2019-04-18 05:48:18")
	time2, _ := time.Parse("2006-01-02 15:04:05", "2019-04-10 05:48:18")
	time3, _ := time.Parse("2006-01-02 15:04:05", "2019-04-8 05:48:18")
	expectedOutput := []time.Time{time3, time2, time1}
	sort.JobByShipment(jobs, sort.ASCENDING)
	for k, v := range expectedOutput {
		if jobs[k].Shipment != v {
			t.Errorf("Sort By Shipment Ascending Failed. Expected : %s ; Actual: %s", v, jobs[k].Shipment)
		}
	}
}

func TestJobSortByShipmentDescending(t *testing.T) {
	jobs := makeDummyJob()
	time1, _ := time.Parse("2006-01-02 15:04:05", "2019-04-18 05:48:18")
	time2, _ := time.Parse("2006-01-02 15:04:05", "2019-04-10 05:48:18")
	time3, _ := time.Parse("2006-01-02 15:04:05", "2019-04-8 05:48:18")
	expectedOutput := []time.Time{time1, time2, time3}
	sort.JobByShipment(jobs, sort.DESCENDING)
	for k, v := range expectedOutput {
		if jobs[k].Shipment != v {
			t.Errorf("Sort By Shipment Descending Failed. Expected : %s ; Actual: %s", v, jobs[k].Shipment)
		}
	}
}

func TestJobSortByDistanceAscending(t *testing.T) {
	jobs := makeDummyJob()
	expectedOutput := []float64{100, 150, 200}
	sort.JobByDistance(jobs, sort.ASCENDING)
	for k, v := range expectedOutput {
		if jobs[k].Distance != v {
			t.Errorf("Sort By Distance Ascending Failed. Expected : %f ; Actual: %f", v, jobs[k].Distance)
		}
	}
}

func TestJobSortByDistanceDescending(t *testing.T) {
	jobs := makeDummyJob()
	expectedOutput := []float64{200, 150, 100}
	sort.JobByDistance(jobs, sort.DESCENDING)
	for k, v := range expectedOutput {
		if jobs[k].Distance != v {
			t.Errorf("Sort By Distance Descending Failed. Expected : %f ; Actual: %f", v, jobs[k].Distance)
		}
	}
}

func TestBidSortByTransporterNameAscending(t *testing.T) {
	bids := makeDummyBid()
	expectedOutput := []string{"Anthon", "Edwin", "Hartoyo"}
	sort.BidByTransporterName(bids, sort.ASCENDING)
	for k, v := range expectedOutput {
		if bids[k].TransporterName != v {
			t.Errorf("Sort By TransporterName Ascending Failed. Expected : %s ; Actual: %s", v, bids[k].TransporterName)
		}
	}
}

func TestBidSortByTransporterNameDescending(t *testing.T) {
	bids := makeDummyBid()
	expectedOutput := []string{"Hartoyo", "Edwin", "Anthon"}
	sort.BidByTransporterName(bids, sort.DESCENDING)
	for k, v := range expectedOutput {
		if bids[k].TransporterName != v {
			t.Errorf("Sort By TransporterName Descending Failed. Expected : %s ; Actual: %s", v, bids[k].TransporterName)
		}
	}
}

func TestBidSortByTransporterRatingAscending(t *testing.T) {
	bids := makeDummyBid()
	expectedOutput := []string{"Bagus", "Baik", "Kurang"}
	sort.BidByTransporterRating(bids, sort.ASCENDING)
	for k, v := range expectedOutput {
		if bids[k].TransporterRating != v {
			t.Errorf("Sort By TransporterRating Ascending Failed. Expected : %s ; Actual: %s", v, bids[k].TransporterRating)
		}
	}
}

func TestBidSortByTransporterRatingDescending(t *testing.T) {
	bids := makeDummyBid()
	expectedOutput := []string{"Kurang", "Baik", "Bagus"}
	sort.BidByTransporterRating(bids, sort.DESCENDING)
	for k, v := range expectedOutput {
		if bids[k].TransporterRating != v {
			t.Errorf("Sort By TransporterRating Descending Failed. Expected : %s ; Actual: %s", v, bids[k].TransporterRating)
		}
	}
}

func TestBidSortByPriceAscending(t *testing.T) {
	bids := makeDummyBid()
	expectedOutput := []float64{2000000, 2100000, 2500000}
	sort.BidByPrice(bids, sort.ASCENDING)
	for k, v := range expectedOutput {
		if bids[k].Price != v {
			t.Errorf("Sort By Price Ascending Failed. Expected : %f ; Actual: %f", v, bids[k].Price)
		}
	}
}

func TestBidSortByPriceDescending(t *testing.T) {
	bids := makeDummyBid()
	expectedOutput := []float64{2500000, 2100000, 2000000}
	sort.BidByPrice(bids, sort.DESCENDING)
	for k, v := range expectedOutput {
		if bids[k].Price != v {
			t.Errorf("Sort By Price Descending Failed. Expected : %f ; Actual: %f", v, bids[k].Price)
		}
	}
}

func TestBidSortByVehicleNameAscending(t *testing.T) {
	bids := makeDummyBid()
	expectedOutput := []string{"Avanza", "Suzuki", "Toyota"}
	sort.BidByVehicleName(bids, sort.ASCENDING)
	for k, v := range expectedOutput {
		if bids[k].VehicleName != v {
			t.Errorf("Sort By VehicleName Ascending Failed. Expected : %s ; Actual: %s", v, bids[k].VehicleName)
		}
	}
}

func TestBidSortByVehicleNameDescending(t *testing.T) {
	bids := makeDummyBid()
	expectedOutput := []string{"Toyota", "Suzuki", "Avanza"}
	sort.BidByVehicleName(bids, sort.DESCENDING)
	for k, v := range expectedOutput {
		if bids[k].VehicleName != v {
			t.Errorf("Sort By VehicleName Descending Failed. Expected : %s ; Actual: %s", v, bids[k].VehicleName)
		}
	}
}
