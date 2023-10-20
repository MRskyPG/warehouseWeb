package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"warehouseWeb/internal/sql"
	"warehouseWeb/pkg/srv/httpserver"
)

func main() {
	var srv httpserver.Server
	fs := http.FileServer(http.Dir("./internal/frontend"))
	go func() {
		if err := srv.Run("80", fs); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	db, err := sql.GetDB()
	if err != nil {
		panic("Didn't to Get Database.")
	}

	rows, err := db.Query("select login from workers")
	if err != nil {
		panic("Error. Database not found")
	}
	for rows.Next() {
		var version string
		_ = rows.Scan(&version)
		fmt.Println(version)
	}

	access, err := sql.GetAccess(db, "vizzcon", "vizzcon")
	if err != nil {
		panic(err)
	}
	if access {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

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
