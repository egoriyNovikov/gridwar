package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type ServerConfig struct {
	Host string
	Port string
}

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		Server: ServerConfig{
			Host: os.Getenv("HOST"),
			Port: os.Getenv("PORT"),
		},
	}

	return cfg, nil
}

func (c *Config) GetDBConfig() *DBConfig {
	return &c.DB
}

func (c *Config) GetServerConfig() *ServerConfig {
	return &c.Server
}

func (c *Config) GetConfig() *Config {
	return c
}
