package repository

import (
	"basic-personal-financial-tracking-api/service/module/domain"

	"gorm.io/gorm"
)

type newRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.PersonalFinancialTrackingRepository {
	return &newRepo{db: db}
}
