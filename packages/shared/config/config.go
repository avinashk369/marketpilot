package config

import (
	"errors"
	"fmt"
)

type Config struct {
	AppEnv string

	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	RedisHost string
	RedisPort string

	JWTSecret string
}

func (c *Config) Validate() error {
	if c.AppEnv == "" {
		return errors.New("APP_ENV is required")
	}

	if c.DBHost == "" {
		return errors.New("DB_HOST is required")
	}

	if c.DBPort == "" {
		return errors.New("DB_PORT is required")
	}

	if c.JWTSecret == "" {
		return errors.New("JWT_SECRET is required")
	}

	return nil
}


func (c *Config) DatabaseURL() string {
    return fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s",
        c.DBUser,
        c.DBPassword,
        c.DBHost,
        c.DBPort,
        c.DBName,
    )
}

func (c *Config) RedisAddress() string {
    return fmt.Sprintf(
        "%s:%s",
        c.RedisHost,
        c.RedisPort,
    )
}