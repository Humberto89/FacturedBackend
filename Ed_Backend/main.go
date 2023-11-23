// main.go
package main

import (
	"Backend/migrations"
	"Backend/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := gin.Default()
	// Configurar la conexión a la base de datos PostgreSQL
	var err error
	db, err := gorm.Open("postgres", "user=postgres password=0000 dbname=DBFormulario sslmode=disable")
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err.Error())
		return
	}
	defer db.Close()

	// Ejecutar migraciones
	migrations.Migrate(db)

	r = routes.SetupRouter(r, db)
	r.Run(":8080")
	fmt.Println("Conexión a la base de datos exitosa.")
}
