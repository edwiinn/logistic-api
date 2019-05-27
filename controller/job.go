package controller

import (
	"logistic-api/core"
	"logistic-api/model/entity"
	repo "logistic-api/model/repository"
	svc "logistic-api/model/service"
	"logistic-api/utils/sort"
	"net/http"
)

type JobController struct {
	core.Controller
	repository *repo.JobRepository
	service    *svc.JobService
}

func (c *JobController) Init() {
	if c.repository == nil {
		c.repository = new(repo.JobRepository)
	}
	if c.service == nil {
		c.service = &svc.JobService{Repository: c.repository}
	}
}
func (c *JobController) GET() {
	jobs := c.service.GetAllJob()
	sortOk := true
	if c.Ctx.QueryParam("sort_by") != "" {
		sortOk = c.processSortRequest(jobs)
	}
	if sortOk {
		c.Ctx.JSON(http.StatusOK, jobs)
	}
}

func (c *JobController) processSortRequest(jobs []entity.Job) bool {
	isOk := true
	sortParam := c.Ctx.QueryParam("sort_by")
	if sortParam != "" {
		orderParams := c.Ctx.QueryParam("order_by")
		if orderParams == "ascending" || orderParams == "asc" || orderParams == "" {
			ok := c.sortJob(jobs, sortParam, sort.ASCENDING)
			if !ok {
				isOk = false
				msg := core.Message{
					StatusCode: http.StatusBadRequest,
					Message:    "sort_by must use one of this parameter : 'origin', 'destination', 'budget', 'shipment', 'distance'"}
				c.Ctx.JSON(http.StatusBadRequest, msg)
			}
		} else if orderParams == "descending" || orderParams == "dsc" {
			ok := c.sortJob(jobs, sortParam, sort.DESCENDING)
			if !ok {
				isOk = false
				msg := core.Message{
					StatusCode: http.StatusBadRequest,
					Message:    "sort_by must use one of this parameter : 'origin', 'destination', 'budget', 'shipment', 'distance'"}
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
func (c *JobController) sortJob(jobs []entity.Job, sortBy string, orderBy sort.OrderBy) bool {
	switch sortBy {
	case "origin":
		sort.JobByOrigin(jobs, orderBy)
	case "destination":
		sort.JobByDestination(jobs, orderBy)
	case "budget":
		sort.JobByBudget(jobs, orderBy)
	case "shipment":
		sort.JobByShipment(jobs, orderBy)
	case "distance":
		sort.JobByDistance(jobs, orderBy)
	default:
		return false
	}
	return true
}
