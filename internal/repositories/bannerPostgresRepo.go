package repositories

import (
	"fmt"
	"gotask/internal/database"
	"gotask/internal/models"

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
