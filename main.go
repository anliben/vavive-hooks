package main

import (
	"anliben/hooks/internal/commons/db"
	"anliben/hooks/pkg/alteracoes"
	"anliben/hooks/pkg/configs"
	"anliben/hooks/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	db, _ := db.OpenConnection()

	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: false,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	alteracoes.RegisterRoutes(app, db)

	utils.StartServerWithGracefulShutdown(app)
}
