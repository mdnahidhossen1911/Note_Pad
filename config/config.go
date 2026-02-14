package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var cfg *Config

type Config struct {
	Version              string
	ServiceName          string
	Port                 string
	JwtSecureKey         string
	JwtExpiryDays        int
	RefreshJwtExpiryDays int
	DbConfig             DBConfig
	AppPass              string
	SenderMail           string
}

func load() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env:", err)
		os.Exit(1)
	}

	required := map[string]*string{
		"VERSION":                 new(string),
		"SERVICE_NAME":            new(string),
		"PORT":                    new(string),
		"JWT_SECURE_KEY":          new(string),
		"JWT_EXPIRY_DAYS":         new(string),
		"REFRESH_JWT_EXPIRY_DAYS": new(string),
		"AppPass":                 new(string),
		"SenderMail":              new(string),
	}

	for key := range required {
		val := os.Getenv(key)
		if val == "" {
			fmt.Printf("%s not set in environment\n", key)
			os.Exit(1)
		}
		*required[key] = val
	}

	jwtExpiryDays, err := strconv.Atoi(os.Getenv("JWT_EXPIRY_DAYS"))
	if err != nil || jwtExpiryDays <= 0 {
		fmt.Println("JWT_EXPIRY_DAYS must be a positive integer")
		os.Exit(1)
	}

	refreshjwtExpiryDays, err := strconv.Atoi(os.Getenv("REFRESH_JWT_EXPIRY_DAYS"))
	if err != nil || refreshjwtExpiryDays <= 0 {
		fmt.Println("REFRESH_JWT_EXPIRY_DAYS must be a positive integer")
		os.Exit(1)
	}

	cfg = &Config{
		Version:              os.Getenv("VERSION"),
		ServiceName:          os.Getenv("SERVICE_NAME"),
		Port:                 os.Getenv("PORT"),
		JwtSecureKey:         os.Getenv("JWT_SECURE_KEY"),
		AppPass:              os.Getenv("AppPass"),
		SenderMail:           os.Getenv("SenderMail"),
		JwtExpiryDays:        jwtExpiryDays,
		RefreshJwtExpiryDays: refreshjwtExpiryDays,
	}

	loadDBConfig()

	fmt.Printf("✅ Config loaded — v%s | port %s\n", cfg.Version, cfg.Port)

}

func GetConfig() *Config {
	if cfg == nil {
		load()
	}
	return cfg
}
