package app

import (
	"context"
	"fmt"
	"gotask/api"
	"gotask/internal/cache"
	"gotask/internal/config"
	"gotask/internal/database"
	"gotask/internal/handlers"
	"gotask/internal/repositories"
	"gotask/internal/services"
	"gotask/sqlc/db_generated"
	"io"
	"time"

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

func (a *App) Stop() error {
	const op = "app.Stop"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.HTTPSrv.Shutdown(ctx); err != nil {
		return fmt.Errorf("%s: failed to shutdown HTTP server: %w", op, err)
	}
	log.Println("HTTP server stopped successfully")

	a.Db.CloseDB()

	log.Println("Database connection closed successfully")

	return nil
}

func (a *App) initBannerHttpHandler() {
	dbPool := a.Db.GetDB()
	queries := db_generated.New(dbPool)

	bannerPostgresRepo := repositories.NewBannerPostgresRepo(queries)
	bannerCache := cache.NewBannerMemoryCache(time.Duration(a.Cfg.Server.CacheTTL * int(time.Minute)))
	bannerServiceImpl := services.NewBannerServiceImpl(bannerPostgresRepo, bannerCache)
	bannerHttpHandler := handlers.NewBannerHttpHandler(bannerServiceImpl)
	bannerApiHandler := handlers.NewBannerApiHandler(bannerServiceImpl)

	bannerRoutes := a.HTTPSrv.Group("v1")
	bannerRoutes.GET("/banners/auction", bannerHttpHandler.GetBannerAuction)
	api.RegisterHandlers(bannerRoutes, bannerApiHandler)
}
