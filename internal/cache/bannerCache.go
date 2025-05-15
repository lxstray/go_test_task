package cache

import (
	"gotask/sqlc/db_generated"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type BannerCache interface {
	Get(geo string, feature int32) (*db_generated.Banner, bool)
	Set(geo string, feature int32, banner *db_generated.Banner, ttl time.Duration)
	Invalidate(id pgtype.UUID)
	InvalidateAll()
}
