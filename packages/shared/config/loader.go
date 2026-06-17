package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	cfg  *Config
	once sync.Once
)

func Load() (*Config, error) {
	var err error

	once.Do(func() {
		_ = godotenv.Load(".env")

		viper.AutomaticEnv()

		cfg = &Config{
			AppEnv: viper.GetString("APP_ENV"),

			DBHost:     viper.GetString("DB_HOST"),
			DBPort:     viper.GetString("DB_PORT"),
			DBName:     viper.GetString("DB_NAME"),
			DBUser:     viper.GetString("DB_USER"),
			DBPassword: viper.GetString("DB_PASSWORD"),

			RedisHost: viper.GetString("REDIS_HOST"),
			RedisPort: viper.GetString("REDIS_PORT"),

			JWTSecret: viper.GetString("JWT_SECRET"),
		}
	})

	return cfg, err
}