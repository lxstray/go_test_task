package repositories

import (
	"context"
	"fmt"
	"gotask/sqlc/db_generated"

	"github.com/jackc/pgx/v5/pgtype"
)

type bannerPostgresRepo struct {
	db *db_generated.Queries
}

func NewBannerPostgresRepo(db *db_generated.Queries) BannerRepo {
	return &bannerPostgresRepo{db: db}
}

func (b *bannerPostgresRepo) SelectTopBanner(ctx context.Context, topBannerParams db_generated.SelectTopBannerParams) (*db_generated.Banner, error) {
	banner, err := b.db.SelectTopBanner(ctx, topBannerParams)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("no banner found for geo=%s and feature=%d", topBannerParams.Geo, topBannerParams.Feature)
		}
		return nil, fmt.Errorf("database error: %w", err)
	}
	return &banner, nil
}

func (b *bannerPostgresRepo) SelectAll(ctx context.Context) (*[]db_generated.Banner, error) {
	banners, err := b.db.SelectAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	if len(banners) == 0 {
		return nil, fmt.Errorf("no banners found")
	}
	return &banners, nil
}

func (b *bannerPostgresRepo) SelectById(ctx context.Context, id pgtype.UUID) (*db_generated.Banner, error) {
	banner, err := b.db.SelectById(ctx, id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("banner with id=%s not found", id.String())
		}
		return nil, fmt.Errorf("database error: %w", err)
	}
	return &banner, nil
}

func (b *bannerPostgresRepo) Create(ctx context.Context, input *db_generated.CreateBannerParams) (pgtype.UUID, error) {
	newBanner, err := b.db.CreateBanner(ctx, *input)
	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("failed to create banner: %w", err)
	}
	return newBanner.ID, nil
}

func (b *bannerPostgresRepo) Update(ctx context.Context, input *db_generated.UpdateBannerParams) error {
	_, err := b.db.UpdateBanner(ctx, *input)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return fmt.Errorf("banner with id=%s not found", input.ID)
		}
		return fmt.Errorf("failed to update banner: %w", err)
	}
	return nil
}

func (b *bannerPostgresRepo) Delete(ctx context.Context, id pgtype.UUID) error {
	err := b.db.DeleteBanner(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete banner: %w", err)
	}
	return nil
}
