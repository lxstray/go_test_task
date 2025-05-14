package repositories

import (
	"fmt"
	"gotask/internal/database"
	"gotask/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type bannerPostgresRepo struct {
	db database.Database
}

func NewBannerPostgresRepo(db database.Database) BannerRepo {
	return &bannerPostgresRepo{db: db}
}

func (b *bannerPostgresRepo) SelectTopBanner(geo string, feature int) (*models.Banner, error) {
	var banner models.Banner

	result := b.db.GetDB().Where("geo = ? AND feature = ?", geo, feature).
		Order("cpm DESC").
		Limit(1).
		First(&banner)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no banner found for geo=%s and feature=%d", geo, feature)
		}
		return nil, fmt.Errorf("database error: %w", result.Error)
	}

	return &banner, nil
}

func (b *bannerPostgresRepo) SelectAll() ([]*models.Banner, error) {
	var banners []*models.Banner

	result := b.db.GetDB().Find(&banners)
	if result.Error != nil {
		return nil, fmt.Errorf("database error: %w", result.Error)
	}

	if len(banners) == 0 {
		return nil, fmt.Errorf("no banners found")
	}

	return banners, nil
}

func (b *bannerPostgresRepo) SelectById(id uuid.UUID) (*models.Banner, error) {
	var banner models.Banner

	result := b.db.GetDB().Where("id = ?", id).First(&banner)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("banner with id=%s not found", id)
		}
		return nil, fmt.Errorf("database error: %w", result.Error)
	}

	return &banner, nil
}

func (b *bannerPostgresRepo) Create(input *models.Banner) error {
	result := b.db.GetDB().Create(input)
	if result.Error != nil {
		return fmt.Errorf("failed to create banner: %w", result.Error)
	}
	return nil
}

func (b *bannerPostgresRepo) Update(id uuid.UUID, input *models.Banner) error {
	var existingBanner models.Banner

	result := b.db.GetDB().Where("id = ?", id).First(&existingBanner)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("banner with id=%s not found", id)
		}
		return fmt.Errorf("database error: %w", result.Error)
	}

	result = b.db.GetDB().Model(&existingBanner).Updates(input)
	if result.Error != nil {
		return fmt.Errorf("failed to update banner: %w", result.Error)
	}

	return nil
}

func (b *bannerPostgresRepo) Delete(id uuid.UUID) error {
	result := b.db.GetDB().Where("id = ?", id).Delete(&models.Banner{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete banner: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("banner with id=%s not found", id)
	}
	return nil
}
