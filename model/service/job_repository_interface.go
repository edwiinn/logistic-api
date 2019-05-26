package service

import (
	"logistic-api/model/entity"
)

type JobRepositoryInterface interface {
	GetAll() []entity.Job
}
