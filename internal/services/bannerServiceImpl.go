package services

import (
	"context"
	"fmt"
	"gotask/internal/cache"
	"gotask/internal/repositories"
	"gotask/sqlc/db_generated"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func (b *bannerServiceImpl) RunBannerAuction(ctx context.Context, geo string, feature int32) (*db_generated.Banner, error) {
	if utf8.RuneCountInString(geo) != 2 {
		return nil, fmt.Errorf("geo must be 2 symbols")
	}

	if feature <= 0 || feature >= 100 {
		return nil, fmt.Errorf("feature must be between 0 and 100")
	}

	if cachedBanner, found := b.cache.Get(geo, feature); found {
		return cachedBanner, nil
	}

	banner, err := b.bannerRepo.SelectTopBanner(ctx, db_generated.SelectTopBannerParams{Geo: geo, Feature: feature})
	if err != nil {
		return nil, err
	}

	if banner != nil {
		b.cache.Set(geo, feature, banner, 0) // 0 это дефолтный TTL
	}

	return banner, nil
}

func (b *bannerServiceImpl) GetAllBanners(ctx context.Context) (*[]db_generated.Banner, error) {
	banners, err := b.bannerRepo.SelectAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all banners: %w", err)
	}
	return banners, nil
}

func (b *bannerServiceImpl) GetBannerById(ctx context.Context, id uuid.UUID) (*db_generated.Banner, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("invalid banner ID")
	}

	banner, err := b.bannerRepo.SelectById(ctx, converToPgtype(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get banner by ID: %w", err)
	}
	return banner, nil
}

func (b *bannerServiceImpl) CreateBanner(ctx context.Context, input *db_generated.CreateBannerParams) error {
	if input == nil {
		return fmt.Errorf("banner input cannot be nil")
	}
	if input.Geo != "" && utf8.RuneCountInString(input.Geo) != 2 {
		return fmt.Errorf("geo must be 2 symbols")
	}
	if input.Feature <= 0 || input.Feature >= 100 {
		return fmt.Errorf("feature must be between 0 and 100")
	}

	id, err := b.bannerRepo.Create(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to create banner: %w", err)
	}

	b.cache.Invalidate(id)
	return nil
}

func (b *bannerServiceImpl) UpdateBanner(ctx context.Context, id uuid.UUID, input *db_generated.CreateBannerParams) error {
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

	err := b.bannerRepo.Update(ctx, &db_generated.UpdateBannerParams{
		ID:      converToPgtype(id),
		Name:    input.Name,
		Image:   input.Image,
		Cpm:     input.Cpm,
		Geo:     input.Geo,
		Feature: input.Feature,
	})
	if err != nil {
		return fmt.Errorf("failed to update banner: %w", err)
	}

	b.cache.Invalidate(converToPgtype(id))
	return nil
}

func (b *bannerServiceImpl) DeleteBanner(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return fmt.Errorf("invalid banner ID")
	}

	err := b.bannerRepo.Delete(ctx, converToPgtype(id))
	if err != nil {
		return fmt.Errorf("failed to delete banner: %w", err)
	}

	b.cache.Invalidate(converToPgtype(id))
	return nil
}

func converToPgtype(u uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: [16]byte(u),
		Valid: true,
	}
}
