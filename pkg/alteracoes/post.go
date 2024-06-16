package alteracoes

import (
	"anliben/hooks/pkg/commons/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) CreateOne(app *fiber.Ctx) error {
	var model models.AlteracaoCampos

	err := app.BodyParser(&model)

	if err != nil {
		err = app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "Alteracao invalida!",
			"error":  err.Error(),
		})
		return err
	}

	err = r.Db.Create(&model).Error

	if err != nil {
		err = app.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"detail": "Alteracao não pode ser criado!",
			"error":  err.Error(),
		})
		return err
	}

	return app.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "alteracao criada com sucesso",
		"item":    model,
	})
}
