package repositories

import (
	"Go_Gin/database"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// ReporteCompra obtiene archivos JSON de MongoDB y extrae datos específicos con paginación.
func ReporteAnexo(c *gin.Context, collections []string) {
	// Conectar a MongoDB
	client, err := database.ConnectdbMongo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al conectar a MongoDB: %v", err)})
		return
	}
	defer client.Disconnect(context.Background())

	// Acceder a la base de datos
	nombreBaseDeDatos := "DTE_Recepcion" // Reemplaza con el nombre de tu base de datos
	database := client.Database(nombreBaseDeDatos)

	// Realizar consultas en todas las colecciones
	var allDocuments []bson.M
	totalRows := int64(0)
	for _, col := range collections {
		if col == "Archivos" {
			continue
		}

		collection := database.Collection(col)

		// Consultar todos los documentos en la colección actual
		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al consultar la colección: %v", err)})
			return
		}
		defer cursor.Close(context.Background())

		// Procesar documentos y extraer datos
		var documents []bson.M
		if err := cursor.All(context.Background(), &documents); err != nil {
			// Manejar error
		}

		allDocuments = append(allDocuments, documents...)

		// Contar documentos en la colección actual y sumar al totalRows
		count, err := collection.CountDocuments(context.Background(), bson.M{})
		if err != nil {
			// Manejar error
		}
		totalRows += count
	}

	// Realizar la paginación sobre todos los documentos acumulados
	pageNumber, _ := strconv.Atoi(c.Query("page"))
	elementsPerPage, _ := strconv.Atoi(c.Query("perPage"))

	startIndex := (pageNumber - 1) * elementsPerPage
	endIndex := startIndex + elementsPerPage

	// Verificar si el índice final no supera el tamaño total
	if endIndex > len(allDocuments) {
		endIndex = len(allDocuments)
	}

	// Obtener la porción de documentos para la página actual
	paginatedDocuments := allDocuments[startIndex:endIndex]

	// Ejemplo de respuesta
	c.JSON(http.StatusOK, gin.H{"archivos": paginatedDocuments, "totalRows": totalRows})
}
