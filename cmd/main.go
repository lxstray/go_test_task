package main

import (
	"gotask/internal/app"
	"gotask/internal/config"
	"gotask/internal/database"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Starting application")

	cfg := config.GetConfig()

	db := database.NewPostgresDB(cfg)

	application := app.New(cfg, db)

	go application.Run()

	//graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Println("Stopping application", sign.String())

	if err := application.Stop(); err != nil {
		panic(err)
	}

	log.Println("Application stopped")
}
