package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (*Handler) GetBannerAuction(c echo.Context) error {
	return c.String(http.StatusOK, "qq")
}
