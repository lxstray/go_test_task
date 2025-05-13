package repositories

import "gotask/internal/models"

type BannerRepo interface {
	SelectTopBanner(geo string, feature int) (*models.Banner, error)
}
