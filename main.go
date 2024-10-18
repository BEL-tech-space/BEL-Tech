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

	app := backend.NewApp(&environmentalVariables, conn)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", environmentalVariables.Port),
		Handler: app.Routes(),
		// ErrorLog:    slog.NewLogLogger(jsonLogger, slog.LevelError),
		// IdleTimeout: time.Minute,
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("Server is running on http://localhost:%v\n", environmentalVariables.Port)
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
