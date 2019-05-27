package service

import (
	"logistic-api/model/entity"
)

type JobService struct {
	Repository JobRepositoryInterface
}

func (s *JobService) GetAll() []entity.Job {
	return s.Repository.GetAll()
}
