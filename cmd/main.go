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

	//TODO: если приложение запускается в горутине(как и должно быть), не выводится ошибки подключения к постгре
	go application.MustRun()

	//graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Println("Stopping application", sign.String())

	//TODO: добавить метод для shutdown если он вообще есть в echo

	log.Println("Application stopped")
}
