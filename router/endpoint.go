package router

const (
	EndpointAPI         = "/api"
	EndpointJobs        = EndpointAPI + "/jobs"
	EndpointBids        = EndpointAPI + "/bids"
	EndpointBidsByJobID = EndpointJobs + "/:job_id/bids"
)
