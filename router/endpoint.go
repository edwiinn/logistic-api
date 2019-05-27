package router

const APIVersion = "1"

const (
	EndpointAPI         = "/api/v" + APIVersion
	EndpointJobs        = EndpointAPI + "/jobs"
	EndpointBids        = EndpointAPI + "/bids"
	EndpointBidsByJobID = EndpointJobs + "/:job_id/bids"
)
