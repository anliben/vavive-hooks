package configs

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

var cfg *config

func FiberConfig() fiber.Config {
	return fiber.Config{
		StreamRequestBody: true,
		ReadTimeout:       90 * time.Second,
		WriteTimeout:      90 * time.Second,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
	}
}

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	Type       string
	Host       string
	Port       string
	User       string
	Pass       string
	Database   string
	Enviroment string
}

func init() {
	cfg = new(config)

	cfg.DB.Type = "postgresql"
	cfg.DB.Database = os.Getenv("PGDATABASE")
	cfg.DB.Host = os.Getenv("PGHOST")
	cfg.DB.Port = os.Getenv("PGPORT")
	cfg.DB.User = os.Getenv("PGUSER")
	cfg.DB.Pass = os.Getenv("PGPASSWORD")
	cfg.DB.Enviroment = os.Getenv("ENVIRONMENT")
}

func GetDB() DBConfig {
	return cfg.DB
}
