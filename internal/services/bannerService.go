package services

import (
	"gotask/internal/models"

	"github.com/google/uuid"
)

type BannerService interface {
	RunBannerAuction(geo string, feature int) (*models.Banner, error)
	GetAllBanners() ([]*models.Banner, error)
	GetBannerById(id uuid.UUID) (*models.Banner, error)
	CreateBanner(input *models.Banner) error
	UpdateBanner(id uuid.UUID, input *models.Banner) error
	DeleteBanner(id uuid.UUID) error
}
