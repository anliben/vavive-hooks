package alteracoes

import (
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	Db *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	r := &handler{
		Db: db,
	}

	routes := app.Group("/api/v1/alteracoes")
	routes.Get("/:id", r.GetAll)
	routes.Post("/one", r.CreateOne)
	routes.Post("/bulk", r.CreateBulk)
}
