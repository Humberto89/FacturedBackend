// migrations/migrate.go
package migrations

import (
	"Backend/models"

	"github.com/jinzhu/gorm"
)

// Migrate ejecuta las migraciones
func Migrate(db *gorm.DB) {

	db.AutoMigrate(
		&models.Pais{},
		&models.Departamento{},
		&models.Municipio{},
		&models.Formulario{},
	)
}
