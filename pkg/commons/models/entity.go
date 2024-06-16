package models

import (
	"gorm.io/gorm"
)

type AlteracaoCampos struct {
	gorm.Model
	CampoAlterado string `json:"campo_alterado"`
	ValorAnterior string `json:"valor_anterior"`
	ValorNovo     string `json:"valor_novo"`
	IDAlteracao   uint   `json:"id_alteracao"`
}
