package service

import (
	"logistic-api/model/entity"
)

type BidRepositoryInterface interface {
	GetAll() []entity.Bid
	GetByJobId(jobID int) []entity.Bid
}
