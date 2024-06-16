package db

import (
	"anliben/hooks/pkg/commons/models"
	"anliben/hooks/pkg/configs"
	"fmt"

	// _ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() (*gorm.DB, error) {

	config := configs.GetDB()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		config.Host,
		config.User,
		config.Pass,
		config.Database,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	// AutoMigrate(db)

	return db, err
}

func AutoMigrate(db *gorm.DB) {

	db.Debug().AutoMigrate(
		&models.AlteracaoCampos{},
	)
}
