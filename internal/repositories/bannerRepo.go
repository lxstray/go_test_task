package repositories

import (
	"gotask/internal/models"

	"github.com/google/uuid"
)

type BannerRepo interface {
	SelectTopBanner(geo string, feature int) (*models.Banner, error)
	SelectAll() ([]*models.Banner, error)
	SelectById(id uuid.UUID) (*models.Banner, error)
	Create(input *models.Banner) error
	Update(id uuid.UUID, input *models.Banner) error
	Delete(id uuid.UUID) error
}
