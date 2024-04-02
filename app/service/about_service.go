package service

import (
	model "dizeto-backend/app/model/about"

	"dizeto-backend/app/repository"

	"github.com/google/uuid"
)

type AboutService interface {
	CreateAbout(title, subtitle, description, note, image string) error
	GetAbout() (*model.About, error)
}

type aboutService struct {
	aboutRepo repository.AboutRepository
}

func NewAboutService(aboutRepo repository.AboutRepository) AboutService {
	return &aboutService{aboutRepo: aboutRepo}
}

func (as *aboutService) CreateAbout(title, subtitle, description, note, image string) error {
	// Generate UUID for about ID
	aboutID := uuid.New()

	// Create new about
	newAbout := &model.About{
		ID:          aboutID,
		Title:       title,
		Subtitle:    subtitle,
		Description: description,
		Note:        note,
		Image:       image,
	}

	// Save new about to repository
	err := as.aboutRepo.CreateAbout(newAbout)
	if err != nil {
		return err
	}

	return nil
}

func (as *aboutService) GetAbout() (*model.About, error) {
	about, err := as.aboutRepo.GetAbout()
	if err != nil {
		return nil, err
	}

	return about, nil
}
