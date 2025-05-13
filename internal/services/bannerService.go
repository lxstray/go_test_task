package services

import "gotask/internal/models"

type BannerService interface {
	RunBannerAuction(geo string, feature int) (*models.Banner, error)
}
