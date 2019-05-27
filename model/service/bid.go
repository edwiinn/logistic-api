package service

import (
	"logistic-api/model/entity"
)

type BidService struct {
	Repository BidRepositoryInterface
}

func (s *BidService) GetAllBid() []entity.Bid {
	return s.Repository.GetAll()
}

func (s *BidService) GetBidByJobID(jobID int) []entity.Bid {
	return s.Repository.GetByJobID(jobID)
}
