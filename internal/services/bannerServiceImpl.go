package services

import (
	"fmt"
	"gotask/internal/models"
	"gotask/internal/repositories"
	"unicode/utf8"
)

type bannerServiceImpl struct {
	bannerRepo repositories.BannerRepo
}

func NewBannerServiceImpl(bannerRepo repositories.BannerRepo) BannerService {
	return &bannerServiceImpl{bannerRepo: bannerRepo}
}

func (b *bannerServiceImpl) RunBannerAuction(geo string, feature int) (*models.Banner, error) {
	if utf8.RuneCountInString(geo) != 2 {
		return nil, fmt.Errorf("geo must be 2 symbols")
	}

	if feature <= 0 || feature >= 100 {
		return nil, fmt.Errorf("feature must be between 0 and 100")
	}

	banner, err := b.bannerRepo.SelectTopBanner(geo, feature)
	if err != nil {
		return nil, err
	}

	return banner, nil
}
