package alteracoes

import (
	"net/http"

	"anliben/hooks/pkg/commons/models"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var model []models.AlteracaoCampos

	id := app.Params("id")

	err := r.Db.
		Where("id_alteracao = ?", id).
		Find(&model).Error

	if err != nil {
		app.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"detail": "Erro ao buscar Campos!",
			"error":  err.Error(),
		})
		return err
	}

	if len(model) == 0 {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"detail": "Nenhum campo encontrado!",
		})
		return nil
	}

	return app.JSON(model)
}
