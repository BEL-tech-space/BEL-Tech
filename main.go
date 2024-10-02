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

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// load environmental variables
	var environmentalVariables backend.EnvConfig
	if err := backend.LoadEnv(&environmentalVariables); err != nil {
		log.Fatalln("Error loading environmental variables: " + err.Error()) // exit
	}

	// connect to database
	conn, err := pgxpool.New(context.Background(), environmentalVariables.Database.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	app, endpoint := backend.NewApp(environmentalVariables, conn), 8090

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", endpoint),
		Handler: app.Routes(environmentalVariables),
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
