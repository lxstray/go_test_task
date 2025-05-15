package repositories

import (
	"context"
	"gotask/sqlc/db_generated"

	"github.com/jackc/pgx/v5/pgtype"
)

type BannerRepo interface {
	SelectTopBanner(ctx context.Context, topBannerParams db_generated.SelectTopBannerParams) (*db_generated.Banner, error)
	SelectAll(ctx context.Context) (*[]db_generated.Banner, error)
	SelectById(ctx context.Context, id pgtype.UUID) (*db_generated.Banner, error)
	Create(ctx context.Context, input *db_generated.CreateBannerParams) (pgtype.UUID, error)
	Update(ctx context.Context, input *db_generated.UpdateBannerParams) error
	Delete(ctx context.Context, id pgtype.UUID) error
}
