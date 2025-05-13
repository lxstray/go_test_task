package app

import (
	"fmt"
	"gotask/internal/config"
	"gotask/internal/database"
	handler "gotask/internal/handlers"
	"io"

	"log"

	"github.com/labstack/echo/v4"
)

type App struct {
	HTTPSrv *echo.Echo
	port    int
}

func New(cfg *config.Config, db database.Database) *App {
	e := echo.New()
	//решил убрать банер и инфу о старте сервера -_-
	e.HideBanner = true
	e.Logger.SetOutput(io.Discard)

	handler := handler.NewHandler()

	e.GET("/banners/auction", handler.GetBannerAuction)

	return &App{
		HTTPSrv: e,
		port:    cfg.Server.Port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "app.Run"

	log.Println("Starting server on port", a.port)

	if err := a.HTTPSrv.Start(fmt.Sprintf(":%d", a.port)); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
