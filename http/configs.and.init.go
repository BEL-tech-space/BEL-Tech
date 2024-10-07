package http

import (
	"beltech/frontend"
	"bytes"
	"fmt"
	"image/png"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var beltechWatermarkLogo []byte

func init() {
	var err error
	beltechWatermarkLogo, err = readPNGAsBytes("assets/images/beltech_logo.png")
	if err != nil {
		log.Fatalln("error reading watermark logo image:", err)
	}
}

func readPNGAsBytes(filePath string) ([]byte, error) {
	imageData, err := frontend.AssetsDir.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded file: %v", err)
	}

	img, err := png.Decode(bytes.NewReader(imageData))
	if err != nil {
		return nil, fmt.Errorf("failed to decode PNG: %v", err)
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return nil, fmt.Errorf("failed to encode PNG: %v", err)
	}

	return buf.Bytes(), nil
}

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
		Url string `env:"DATABASE_URL"`
	}

	Port string `env:"PORT"`

	CookieSecret  string `env:"COOKIE_SECRET"`
	SessionSecret string `env:"SESSION_SECRET"`
	Build         string `env:"BUILD"`
}

func LoadEnv(cfg *EnvConfig) error {

	if err := godotenv.Load("prod.env"); err != nil {
		fmt.Printf("error loading prod.env file: %v", err)
	}

	if err := env.Parse(cfg); err != nil {
		return err
	}
	return nil
}
