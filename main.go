package main

import (
	"Go_Gin/database"
	"Go_Gin/migrations"
	"Go_Gin/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//Inicializacion de conexcion a postgre
	db, err := database.ConnectdbPostgre()
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err.Error())
		return
	}
	defer db.Close()

	// Ejecutar migraciones
	migrations.Migrate(db)

	r := gin.Default()

	r = routes.SetupRouter(r, db)
	r.Run(":8080")
	fmt.Println("Conexión a la base de datos exitosa.")

}
