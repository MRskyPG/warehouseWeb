package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"warehouseWeb/pkg/srv/httpserver"
)

func main() {
	var srv httpserver.Server
	fs := http.FileServer(http.Dir("./internal/frontend"))
	//Using graceful shutdown
	go func() {
		if err := srv.Run("80", fs); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Web-app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Web-app Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error was occured on server shutting down: %s", err.Error())
	}
}
