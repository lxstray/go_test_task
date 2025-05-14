package services

import (
	"fmt"
	"gotask/internal/cache"
	"gotask/internal/models"
	"gotask/internal/repositories"
	"unicode/utf8"

	"github.com/google/uuid"
)

type bannerServiceImpl struct {
	bannerRepo repositories.BannerRepo
	cache      cache.BannerCache
}

func NewBannerServiceImpl(bannerRepo repositories.BannerRepo, cache cache.BannerCache) BannerService {
	return &bannerServiceImpl{
		bannerRepo: bannerRepo,
		cache:      cache,
	}
}

func (b *bannerServiceImpl) RunBannerAuction(geo string, feature int) (*models.Banner, error) {
	if utf8.RuneCountInString(geo) != 2 {
		return nil, fmt.Errorf("geo must be 2 symbols")
	}

	if feature <= 0 || feature >= 100 {
		return nil, fmt.Errorf("feature must be between 0 and 100")
	}

	if cachedBanner, found := b.cache.Get(geo, feature); found {
		return cachedBanner, nil
	}

	banner, err := b.bannerRepo.SelectTopBanner(geo, feature)
	if err != nil {
		return nil, err
	}

	if banner != nil {
		b.cache.Set(geo, feature, banner, 0) // 0 это дефолтный TTL
	}

	return banner, nil
}

func (b *bannerServiceImpl) GetAllBanners() ([]*models.Banner, error) {
	banners, err := b.bannerRepo.SelectAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all banners: %w", err)
	}
	return banners, nil
}

func (b *bannerServiceImpl) GetBannerById(id uuid.UUID) (*models.Banner, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("invalid banner ID")
	}

	banner, err := b.bannerRepo.SelectById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get banner by ID: %w", err)
	}
	return banner, nil
}

func (b *bannerServiceImpl) CreateBanner(input *models.Banner) error {
	if input == nil {
		return fmt.Errorf("banner input cannot be nil")
	}
	if input.Geo != "" && utf8.RuneCountInString(input.Geo) != 2 {
		return fmt.Errorf("geo must be 2 symbols")
	}
	if input.Feature <= 0 || input.Feature >= 100 {
		return fmt.Errorf("feature must be between 0 and 100")
	}
	if input.CPM < 0 {
		return fmt.Errorf("CPM must be non-negative")
	}

	err := b.bannerRepo.Create(input)
	if err != nil {
		return fmt.Errorf("failed to create banner: %w", err)
	}

	b.cache.Invalidate(input.ID)
	return nil
}

func (b *bannerServiceImpl) UpdateBanner(id uuid.UUID, input *models.Banner) error {
	if id == uuid.Nil {
		return fmt.Errorf("invalid banner ID")
	}
	if input == nil {
		return fmt.Errorf("banner input cannot be nil")
	}
	if input.Geo != "" && utf8.RuneCountInString(input.Geo) != 2 {
		return fmt.Errorf("geo must be 2 symbols")
	}
	if input.Feature <= 0 || input.Feature >= 100 {
		return fmt.Errorf("feature must be between 0 and 100")
	}
	if input.CPM < 0 {
		return fmt.Errorf("CPM must be non-negative")
	}

	err := b.bannerRepo.Update(id, input)
	if err != nil {
		return fmt.Errorf("failed to update banner: %w", err)
	}

	b.cache.Invalidate(id)
	return nil
}

func (b *bannerServiceImpl) DeleteBanner(id uuid.UUID) error {
	if id == uuid.Nil {
		return fmt.Errorf("invalid banner ID")
	}

	err := b.bannerRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete banner: %w", err)
	}

	b.cache.Invalidate(id)
	return nil
}
