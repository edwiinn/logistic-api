package main

import (
	"encoding/json"
	"logistic-api/core"
	"logistic-api/model/entity"
	"logistic-api/router"
	"logistic-api/utils/sort"
	"net/http"
	"strconv"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const URL = "http://localhost:8080"

func TestGetJobs(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointJobs, nil)
	data := []entity.Job{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBids(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointBids, nil)
	data := []entity.Bid{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBidsByJobId(t *testing.T) {
	go core.Run()
	jobID := 1
	jobIDString := strconv.Itoa(jobID)
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointJobs+"/"+jobIDString+"/bids", nil)
	data := []entity.Bid{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
	isOk := true
	for _, bid := range data {
		if bid.JobID != jobID {
			isOk = false
		}
	}
	if !isOk {
		t.Error("Diffrent Job ID in Bids")
	}
}

func TestGetJobsByOrigin(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointJobs+"?sort_by=origin", nil)
	data := []entity.Job{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
	if !sort.IsSorted(sort.JobByOriginProperties(data)) {
		t.Error("Not Sorted By Origin")
	}
}

func TestGetJobsByDestination(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointJobs+"?sort_by=destination", nil)
	data := []entity.Job{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
	if !sort.IsSorted(sort.JobByDestinationProperties(data)) {
		t.Error("Not Sorted By Destination")
	}
}

func TestGetJobsByBudget(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointJobs+"?sort_by=budget", nil)
	data := []entity.Job{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
	if !sort.IsSorted(sort.JobByBudgetProperties(data)) {
		t.Error("Not Sorted By Budget")
	}
}

func TestGetJobsByShipment(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointJobs+"?sort_by=shipment", nil)
	data := []entity.Job{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
	if !sort.IsSorted(sort.JobByShipmentProperties(data)) {
		t.Error("Not Sorted By Shipment")
	}
}

func TestGetJobsByDistance(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointJobs+"?sort_by=distance", nil)
	data := []entity.Job{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
	if !sort.IsSorted(sort.JobByDistanceProperties(data)) {
		t.Error("Not Sorted By Distance")
	}
}

func TestGetBidsByTransporterName(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointBids+"?sort_by=transporter-name", nil)
	data := []entity.Bid{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
	if !sort.IsSorted(sort.BidByTransporterNameProperties(data)) {
		t.Error("Not Sorted By Transporter Name")
	}
}

func TestGetBidsByTransporterRating(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointBids+"?sort_by=transporter-rating", nil)
	data := []entity.Bid{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
	if !sort.IsSorted(sort.BidByTransporterRatingProperties(data)) {
		t.Error("Not Sorted By Transporter Rating")
	}
}

func TestGetBidsByPrice(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointBids+"?sort_by=price", nil)
	data := []entity.Bid{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
	if !sort.IsSorted(sort.BidByPriceProperties(data)) {
		t.Error("Not Sorted By Price")
	}
}

func TestGetBidsByVehicle(t *testing.T) {
	go core.Run()
	req, err := http.NewRequest(http.MethodGet, URL+router.EndpointBids+"?sort_by=vehicle", nil)
	data := []entity.Bid{}
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Error(err)
	}
	if !sort.IsSorted(sort.BidByVehicleNameProperties(data)) {
		t.Error("Not Sorted By Vehicle")
	}
}
