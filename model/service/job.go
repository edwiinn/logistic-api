package service

import (
	"logistic-api/model/entity"
)

type JobService struct {
	Repository JobRepositoryInterface
}

func (s *JobService) GetAllJob() []entity.Job {
	return s.Repository.GetAll()
}
