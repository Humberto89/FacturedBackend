package repositories

import (
	"Go_Gin/database"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ObtenerArchivosYDatos obtiene archivos JSON de MongoDB y extrae datos específicos con paginación.
func ReporteCompra(c *gin.Context) {
	// Conectar a MongoDB
	client, err := database.ConnectdbMongo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al conectar a MongoDB: %v", err)})
		return
	}
	defer client.Disconnect(context.Background())

	// Acceder a la base de datos
	nombreBaseDeDatos := "ArchivosPrueba" // Reemplaza con el nombre de tu base de datos
	database := client.Database(nombreBaseDeDatos)

	collections := []string{
		"Factura_de_exportacion",
		"Nota_de_debito",
		"Nota_de_credito",
		"Comprobante_de_credito_fiscal",
	}

	// Realizar consultas en colecciones específicas
	var archivos []bson.M
	for _, col := range collections {
		// Excluir colección específica
		if col == "Archivos" {
			continue
		}

		collection := database.Collection(col)

		// Configurar índice de inicio y límite para paginación
		pageNumber, _ := strconv.Atoi(c.Query("page"))
		elementsPerPage, _ := strconv.Atoi(c.Query("perPage"))
		startIndex := (pageNumber - 1) * elementsPerPage
		limit := int64(elementsPerPage)

		options := options.Find()
		if startIndex != 0 {
			options.SetSkip(int64(startIndex))
		}

		if limit != 0 {
			options.SetLimit(limit)
		}

		// Consultar archivos en la colección actual
		cursor, err := collection.Find(context.Background(), bson.M{}, options)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al consultar archivos en la colección %s: %v", col, err)})
			return
		}
		defer cursor.Close(context.Background())

		// Procesar archivos y extraer datos
		var archivosColeccion []bson.M
		if err := cursor.All(context.Background(), &archivosColeccion); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al procesar archivos en la colección %s: %v", col, err)})
			return
		}

		archivos = append(archivos, archivosColeccion...)
	}

	// Puedes procesar la lista de archivos aquí según tus necesidades

	// Ejemplo de respuesta
	c.JSON(http.StatusOK, gin.H{"archivos": archivos})
}
