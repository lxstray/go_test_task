package handlers

import "github.com/labstack/echo/v4"

type BannerHanderl interface {
	GetBannerAuction(c echo.Context) error
}
