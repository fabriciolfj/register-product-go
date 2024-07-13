package main

import (
	"context"
	log "github.com/fabriciolfj/register-product-go/configuration"
	"github.com/fabriciolfj/register-product-go/controller"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	mux := http.NewServeMux()

	pc, _ := InitializeProductControllerWire()
	http.HandleFunc("/products", controller.RecoveryMiddleware(pc.HandleProduct))

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Log.Fatal("fail star server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Log.Fatal("Server forced to shutdown: ", err)
	}

	log.Log.Info("Server exiting")
}
