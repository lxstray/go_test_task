package services

import (
	"context"
	"gotask/sqlc/db_generated"

	"github.com/google/uuid"
)

type BannerService interface {
	RunBannerAuction(ctx context.Context, geo string, feature int32) (*db_generated.Banner, error)
	GetAllBanners(ctx context.Context) (*[]db_generated.Banner, error)
	GetBannerById(ctx context.Context, id uuid.UUID) (*db_generated.Banner, error)
	CreateBanner(ctx context.Context, input *db_generated.CreateBannerParams) error
	UpdateBanner(ctx context.Context, id uuid.UUID, input *db_generated.CreateBannerParams) error
	DeleteBanner(ctx context.Context, id uuid.UUID) error
}
