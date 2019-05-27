package router

import (
	"logistic-api/controller"
	"logistic-api/core"
	"logistic-api/core/request"
)

func init() {
	core.Route(EndpointBids, new(controller.BidController), request.GET)
}
