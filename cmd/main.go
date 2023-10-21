package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"warehouseWeb/internal/handler"
	"warehouseWeb/pkg/srv/httpserver"
)

func main() {
	//Initialize server with main handler
	var srv httpserver.Server
	http.HandleFunc("/", handler.Handler)
	//fs := http.FileServer(http.Dir("./internal/frontend"))

	go func() {
		if err := srv.Run("80", nil); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	//Graceful Shutdown
	logrus.Print("Web-app Started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Web-app Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error was occured on server shutting down: %s", err.Error())
	}

}
