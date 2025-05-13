package handlers

import "github.com/labstack/echo/v4"

type BannerHandler interface {
	GetBannerAuction(c echo.Context) error
}
