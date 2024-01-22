package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/harshvsinghme/uniblox-assmt-backend/dbUtils"
	"github.com/harshvsinghme/uniblox-assmt-backend/global"
	"github.com/harshvsinghme/uniblox-assmt-backend/router"
)

func init() {
	global.LoadGlobals(".env")

	dbUtils.InitDB()
}

func main() {

	appRouter := router.InitRouter()
	srv := &http.Server{
		Addr:    ":3001",
		Handler: appRouter,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("server error encountered", err)
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Println("server shutdown:", err)
	} else {
		log.Println("server healthy shutdown")
	}
}
