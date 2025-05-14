package handlers

import (
	"gotask/api"
	"gotask/internal/models"
	"gotask/internal/services"
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
	banners, err := b.bannerService.GetAllBanners()
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, banners)
	return nil
}

func (b *bannerApiHandler) CreateBanner(c echo.Context) error {
	var newBanner models.Banner
	if err := c.Bind(&newBanner); err != nil {
		return err
	}

	if err := b.bannerService.CreateBanner(&newBanner); err != nil {
		return err
	}

	c.String(http.StatusOK, "banner created")
	return nil
}

// TODO: работа над ошибками
func (b *bannerApiHandler) DeleteBanner(c echo.Context, id openapi_types.UUID) error {
	if err := b.bannerService.DeleteBanner(id); err != nil {
		return err
	}

	c.String(http.StatusOK, "banner deleted")
	return nil
}

func (b *bannerApiHandler) GetBannerById(c echo.Context, id openapi_types.UUID) error {
	banner, err := b.bannerService.GetBannerById(id)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, banner)
	return nil
}

// TODO: добавить проверок нового баннера
func (b *bannerApiHandler) UpdateBanner(c echo.Context, id openapi_types.UUID) error {
	var newBanner models.Banner
	if err := c.Bind(&newBanner); err != nil {
		return err
	}

	if err := b.bannerService.UpdateBanner(id, &newBanner); err != nil {
		return err
	}

	c.String(http.StatusOK, "banner updated")
	return nil
}
