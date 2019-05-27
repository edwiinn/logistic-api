package service

import (
	"logistic-api/model/entity"
)

type BidRepositoryInterface interface {
	GetAll() []entity.Bid
	GetByJobID(jobID int) []entity.Bid
}
