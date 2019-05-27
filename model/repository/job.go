package repository

import (
	"logistic-api/config"
	"logistic-api/model/entity"

	"github.com/jinzhu/gorm"
)

type JobRepository struct{}

// GetAll() []entity.Job
func (r *JobRepository) GetAll() []entity.Job {
	db, err := gorm.Open(config.DatabaseDriver, config.DatabaseSource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	jobs := []entity.Job{}
	db.Find(jobs)
	return jobs
}
