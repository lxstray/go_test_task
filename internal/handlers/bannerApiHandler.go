package handlers

import (
	"fmt"
	"gotask/api"
	"gotask/internal/services"
	"gotask/sqlc/db_generated"
	"net/http"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type bannerApiHandler struct {
	bannerService services.BannerService
}

func NewBannerApiHandler(bannerService services.BannerService) api.ServerInterface {
	return &bannerApiHandler{bannerService: bannerService}
}

func (b *bannerApiHandler) GetBanners(c echo.Context) error {
	banners, err := b.bannerService.GetAllBanners(c.Request().Context())
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
		return fmt.Errorf("failed to get banners: %w", err)
	}

	c.JSON(http.StatusOK, banners)
	return nil
}

func (b *bannerApiHandler) CreateBanner(c echo.Context) error {
	var newBanner db_generated.CreateBannerParams
	if err := c.Bind(&newBanner); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid request body: %s", err))
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	if err := b.bannerService.CreateBanner(c.Request().Context(), &newBanner); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
		return fmt.Errorf("failed to create banner: %w", err)
	}

	c.String(http.StatusOK, "banner created")
	return nil
}

func (b *bannerApiHandler) DeleteBanner(c echo.Context, id openapi_types.UUID) error {
	if err := b.bannerService.DeleteBanner(c.Request().Context(), id); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
		return fmt.Errorf("failed to delete banner %s: %w", id, err)
	}

	c.String(http.StatusOK, "banner deleted")
	return nil
}

func (b *bannerApiHandler) GetBannerById(c echo.Context, id openapi_types.UUID) error {
	banner, err := b.bannerService.GetBannerById(c.Request().Context(), id)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
		return fmt.Errorf("failed to get banner %s: %w", id, err)
	}

	c.JSON(http.StatusOK, banner)
	return nil
}

func (b *bannerApiHandler) UpdateBanner(c echo.Context, id openapi_types.UUID) error {
	var newBanner db_generated.CreateBannerParams
	if err := c.Bind(&newBanner); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid request body: %s", err))
		return fmt.Errorf("failed to bind request body: %w", err)
	}

	if err := b.bannerService.UpdateBanner(c.Request().Context(), id, &newBanner); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
		return fmt.Errorf("failed to update banner %s: %w", id, err)
	}

	c.String(http.StatusOK, "banner updated")
	return nil
}
