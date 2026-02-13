package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func loadDBConfig() {
	fields := map[string]string{
		"DB_HOST":     "",
		"DB_PORT":     "",
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_NAME":     "",
		"DB_SSLMODE":  "",
	}
	for key := range fields {
		val := os.Getenv(key)
		if val == "" {
			fmt.Printf("%s not set in environment\n", key)
			os.Exit(1)
		}
		fields[key] = val
	}

	cfg.DbConfig = DBConfig{
		Host:     fields["DB_HOST"],
		Port:     fields["DB_PORT"],
		User:     fields["DB_USER"],
		Password: fields["DB_PASSWORD"],
		DBName:   fields["DB_NAME"],
		SSLMode:  fields["DB_SSLMODE"],
	}
}
