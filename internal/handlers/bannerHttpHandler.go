package handlers

import (
	"fmt"
	"gotask/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type bannerHttpHandler struct {
	bannerService services.BannerService
}

func NewBannerHttpHandler(bannerService services.BannerService) BannerHandler {
	return &bannerHttpHandler{bannerService: bannerService}
}

//TODO: переделать работу с ошибками(отправлять их клиенту)

func (b *bannerHttpHandler) GetBannerAuction(c echo.Context) error {
	geo := c.QueryParam("geo")
	if geo == "" {
		return fmt.Errorf("missing required parameter: geo")
	}
	feature := c.QueryParam("feature")
	if feature == "" {
		return fmt.Errorf("missing required parameter: feature")
	}

	intFeature, err := strconv.Atoi(feature)
	if err != nil {
		return err
	}

	banner, err := b.bannerService.RunBannerAuction(geo, intFeature)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, banner)

	return nil
}
