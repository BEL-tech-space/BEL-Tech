package main

import (
	backend "beltech/http"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	app, endpoint := backend.NewApp(), 8090

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", endpoint),
		Handler: app.Routes(),
	}

	go func() {
		log.Printf("Server is running on http://localhost:%d/signin \n", endpoint)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panic("Server error: " + err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	<-done
	close(done)

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Panic("Graceful server shutdown Failed: " + err.Error())
	}

	log.Println("SERVER STOPPED GRACEFULLY")
}
