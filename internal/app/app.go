package app

import (
	"fmt"
	"gotask/internal/config"
	"gotask/internal/database"
	"gotask/internal/handlers"
	"gotask/internal/repositories"
	"gotask/internal/services"
	"io"

	"log"

	"github.com/labstack/echo/v4"
)

type App struct {
	HTTPSrv *echo.Echo
	Db      database.Database
	Cfg     *config.Config
}

func New(cfg *config.Config, db database.Database) *App {
	e := echo.New()
	//решил убрать банер и инфу о старте сервера -_-
	e.HideBanner = true
	e.Logger.SetOutput(io.Discard)

	return &App{
		HTTPSrv: e,
		Db:      db,
		Cfg:     cfg,
	}
}

func (a *App) Run() {
	const op = "app.Run"

	a.initBannerHttpHandler()
	log.Println("Starting server on port", a.Cfg.Server.Port)

	if err := a.HTTPSrv.Start(fmt.Sprintf(":%d", a.Cfg.Server.Port)); err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

}

func (a *App) initBannerHttpHandler() {
	bannerPostgresRepo := repositories.NewBannerPostgresRepo(a.Db)
	bannerServiceImpl := services.NewBannerServiceImpl(bannerPostgresRepo)
	bannerHttpHandler := handlers.NewBannerHttpHandler(bannerServiceImpl)

	bannerRoutes := a.HTTPSrv.Group("v1/banners")
	bannerRoutes.GET("/auction", bannerHttpHandler.GetBannerAuction)
}
