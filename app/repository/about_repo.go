package repository

import (
	model "dizeto-backend/app/model/about"

	"github.com/jinzhu/gorm"
)

type AboutRepository interface {
	CreateAbout(about *model.About) error
	GetAbout() (*model.About, error)
}

type aboutRepository struct {
	db *gorm.DB
}

func NewAboutRepository(db *gorm.DB) AboutRepository {
	return &aboutRepository{db: db}
}

func (ar *aboutRepository) CreateAbout(about *model.About) error {
	if err := about.Validate(); err != nil {
		return err
	}
	ar.db.Create(about)
	return nil
}

func (ar *aboutRepository) GetAbout() (*model.About, error) {
	var about model.About
	err := ar.db.First(&about).Error
	return &about, err
}
