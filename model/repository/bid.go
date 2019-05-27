package repository

import (
	"logistic-api/config"
	"logistic-api/model/entity"

	"github.com/jinzhu/gorm"
)

type BidRepository struct{}

func (r *BidRepository) GetAll() []entity.Bid {
	db, err := gorm.Open(config.DatabaseDriver, config.DatabaseSource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	bids := []entity.Bid{}
	db.Find(&bids)
	return bids
}

func (r *BidRepository) GetByJobID(jobID int) []entity.Bid {
	db, err := gorm.Open(config.DatabaseDriver, config.DatabaseSource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	bids := []entity.Bid{}
	db.Where("job_id = ?", jobID).Find(&bids)
	return bids
}
