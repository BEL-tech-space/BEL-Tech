package utils

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	LoginProvider struct {
		Google struct {
			ClientKey string `env:"CLIENT_PROVIDER_KEY_GOOGLE"`
			Secret    string `env:"CLIENT_PROVIDER_SECRET_GOOGLE"`
		}
		Facebook struct {
			ClientKey string `env:"CLIENT_PROVIDER_KEY_FACEBOOK"`
			Secret    string `env:"CLIENT_PROVIDER_SECRET_FACEBOOK"`
		}
	}
	Database struct {
	}

	CookieSecret  string `env:"COOKIE_SECRET"`
	SessionSecret string `env:"SESSION_SECRET"`
	Build         string `env:"BUILD"`
}

func LoadEnv(cfg *EnvConfig) error {

	if err := godotenv.Load("prod.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := env.Parse(cfg); err != nil {
		return err
	}
	fmt.Printf("\n%+v\n\n", cfg)
	return nil
}
