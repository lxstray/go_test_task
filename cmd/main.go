package main

import (
	"gotask/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Starting application")

	//TODO: вынести порт и тд в конфиг
	application := app.New(8080)

	go application.MustRun()

	//graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Println("Stopping application", sign.String())

	//TODO: добавить метод для shutdown если он вообще есть в echo

	log.Println("Application stopped")
}
