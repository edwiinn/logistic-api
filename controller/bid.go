package controller

import (
	"logistic-api/core"
	"logistic-api/model/entity"
	repo "logistic-api/model/repository"
	svc "logistic-api/model/service"
	"logistic-api/utils/sort"
	"net/http"
	"strconv"
)

type BidController struct {
	core.Controller
	repository *repo.BidRepository
	service    *svc.BidService
}

func (c *BidController) Init() {
	if c.repository == nil {
		c.repository = new(repo.BidRepository)
	}
	if c.service == nil {
		c.service = &svc.BidService{Repository: c.repository}
	}
}
func (c *BidController) GET() {
	bids := []entity.Bid{}
	isOk := true
	if c.Ctx.Param("job_id") == "" {
		bids = c.service.GetAllBid()
	} else {
		param, err := strconv.Atoi(c.Ctx.Param("job_id"))
		if err != nil {
			isOk = false
			msg := core.Message{
				StatusCode: http.StatusBadRequest,
				Message:    "param must be only number"}
			c.Ctx.JSON(http.StatusBadRequest, msg)
		}
		bids = c.service.GetBidByJobID(param)
	}
	if c.Ctx.QueryParam("sort_by") != "" {
		isOk = c.processSortRequest(bids)
	}
	if isOk {
		c.Ctx.JSON(http.StatusOK, bids)
	}
}

func (c *BidController) processSortRequest(bids []entity.Bid) bool {
	isOk := true
	sortParam := c.Ctx.QueryParam("sort_by")
	if sortParam != "" {
		orderParams := c.Ctx.QueryParam("order_by")
		if orderParams == "ascending" || orderParams == "asc" || orderParams == "" {
			ok := c.sortBid(bids, sortParam, sort.ASCENDING)
			if !ok {
				isOk = false
				msg := core.Message{
					StatusCode: http.StatusBadRequest,
					Message:    "sort_by must use one of this parameter : 'transporter-name', 'transporter-rating', 'price', 'vehicle'"}
				c.Ctx.JSON(http.StatusBadRequest, msg)
			}
		} else if orderParams == "descending" || orderParams == "dsc" {
			ok := c.sortBid(bids, sortParam, sort.DESCENDING)
			if !ok {
				isOk = false
				msg := core.Message{
					StatusCode: http.StatusBadRequest,
					Message:    "sort_by must use one of this parameter : 'transporter-name', 'transporter-rating', 'price', 'vehicle'"}
				c.Ctx.JSON(http.StatusBadRequest, msg)
			}
		} else {
			isOk = false
			msg := core.Message{
				StatusCode: http.StatusBadRequest,
				Message:    "order_by must use one of this parameter : 'ascending', 'asc', 'descending', 'dsc'"}
			c.Ctx.JSON(http.StatusBadRequest, msg)
		}
	}
	return isOk
}
func (c *BidController) sortBid(bids []entity.Bid, sortBy string, orderBy sort.OrderBy) bool {
	switch sortBy {
	case "transporter-name":
		sort.BidByTransporterName(bids, orderBy)
	case "transporter-rating":
		sort.BidByTransporterRating(bids, orderBy)
	case "price":
		sort.BidByPrice(bids, orderBy)
	case "vehicle":
		sort.BidByVehicleName(bids, orderBy)
	default:
		return false
	}
	return true
}
