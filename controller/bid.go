package controller

import (
	"logistic-api/core"
	"logistic-api/model/service"
)

type BidController struct {
	core.Controller
	service service.BidService
}
