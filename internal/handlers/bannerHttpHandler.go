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

func (b *bannerHttpHandler) GetBannerAuction(c echo.Context) error {
	geo := c.QueryParam("geo")
	if geo == "" {
		c.String(http.StatusBadRequest, "missing required parameter: geo")
		return fmt.Errorf("missing required parameter: geo")
	}
	feature := c.QueryParam("feature")
	if feature == "" {
		c.String(http.StatusBadRequest, "missing required parameter: feature")
		return fmt.Errorf("missing required parameter: feature")
	}

	intFeature, err := strconv.Atoi(feature)
	if err != nil {
		return err
	}

	banner, err := b.bannerService.RunBannerAuction(geo, intFeature)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
		return err
	}

	c.JSON(http.StatusOK, banner)

	return nil
}
