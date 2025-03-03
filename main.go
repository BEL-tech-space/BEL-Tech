package main

import (
	backend "beltech/http"
	"beltech/service"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dgraph-io/ristretto/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "golang.org/x/crypto/x509roots/fallback" // CA bundle for FROM Scratch
)

func main() {
	// load environmental variables
	var environmentalVariables backend.EnvConfig
	if err := backend.LoadEnv(&environmentalVariables); err != nil {
		log.Fatalln("Error loading environmental variables: " + err.Error()) // exit
	}

	// connect to database
	var databaseUrl string
	if environmentalVariables.Build == "production" {
		databaseUrl = environmentalVariables.Database.ProdPrivate
	} else {
		databaseUrl = environmentalVariables.Database.Dev
	}
	conn, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create db connection pool: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	accountCache, err := buildCache[service.Account]()
	if err != nil {
		log.Fatalln("Error creating cache: " + err.Error())
	}
	defer accountCache.Close()
	wasteLogCache, err := buildCache[service.WasteLogFromCollector]()
	if err != nil {
		log.Fatalln("Error creating cache: " + err.Error())
	}
	defer wasteLogCache.Close()

	app := backend.NewApp(&environmentalVariables, conn, accountCache, wasteLogCache)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", environmentalVariables.Port),
		Handler: app.Routes(),
		// ErrorLog:    slog.NewLogLogger(jsonLogger, slog.LevelError),
		// IdleTimeout: time.Minute,
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	go func() {
		if environmentalVariables.Build != "production" {
			log.Printf("Server is running on http://localhost:%v\n", environmentalVariables.Port)
		} else {
			fmt.Println("running in production build")
		}

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panic("Server error: " + err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	stopScrapping := app.ScrapeCryptoData()
	<-done
	close(done)

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Fatalln("Graceful server shutdown Failed: " + err.Error())
	}

	stopScrapping()

	log.Println("SERVER STOPPED GRACEFULLY")
}

func buildCache[T any]() (*ristretto.Cache[string, T], error) {
	cache, err := ristretto.NewCache(&ristretto.Config[string, T]{
		NumCounters: 1e6,     // number of keys to track frequency of (1M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	if err != nil {
		return nil, err
	}

	return cache, nil
}
