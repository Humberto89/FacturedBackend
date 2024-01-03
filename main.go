package main

import (
	"Go_Gin/database"
	"Go_Gin/migrations"
	"Go_Gin/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Inicializacion de conexcion a postgre
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar archivo .env")
	}
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
	r.Run(":8081")
	fmt.Println("Conexión a la base de datos exitosa.")

}
