package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type LogFileConfig struct {
	FilePath string
}

type CsvFileConfig struct {
	FilePath string
}

type TokenConfig struct {
	IssuerName      string
	JwtSignatureKey []byte
	JwtLifeTime     time.Duration
}

type Config struct {
	ApiConfig
	DbConfig
	LogFileConfig
	CsvFileConfig
	TokenConfig
}

const BaseURL = "https://108.136.239.242"
const ImageUploadDirectory = "./uploads"

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.LogFileConfig = LogFileConfig{FilePath: os.Getenv("LOG_FILE")}
	c.CsvFileConfig = CsvFileConfig{FilePath: os.Getenv("CSV_FILE")}

	tokenLifeTime, err := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))
	if err != nil {
		return err
	}

	c.TokenConfig = TokenConfig{
		IssuerName:      os.Getenv("TOKEN_ISSUE_NAME"),
		JwtSignatureKey: []byte(os.Getenv("TOKEN_KEY")),
		JwtLifeTime:     time.Duration(tokenLifeTime) * time.Hour,
	}

	if c.ApiConfig.ApiPort == "" || c.DbConfig.Driver == "" || c.DbConfig.Host == "" || c.DbConfig.Name == "" || c.DbConfig.Port == "" || c.DbConfig.User == "" {
		return errors.New("all environment variables required")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}
