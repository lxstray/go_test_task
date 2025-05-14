package cache

import (
	"gotask/internal/models"
	"time"

	"github.com/google/uuid"
)

type BannerCache interface {
	Get(geo string, feature int) (*models.Banner, bool)
	Set(geo string, feature int, banner *models.Banner, ttl time.Duration)
	Invalidate(id uuid.UUID)
	InvalidateAll()
}
