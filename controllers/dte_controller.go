// controllers/dte_controller.go
package controllers

import (
	"Go_Gin/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"
)

// Controlador para obtener DTEs según parámetros de la URL
func GetDTEs(c *gin.Context, db *gorm.DB) {

	// Obtener parámetros de la URL
	tipoDTEParam := c.Query("tipoDTE")
	fechaInicioParam := c.Query("fechaInicio")
	fechaFinParam := c.Query("fechaFin")
	condicionOperacionParam := c.Query("condicionOperacion")

	// Construir filtro para la consulta
	filter := bson.M{}
	if fechaInicioParam != "" && fechaFinParam != "" {
		// Usar notación de puntos para filtrar por fecha dentro del objeto identificacion
		filter["data.identificacion.fecEmi"] = bson.M{
			"$gte": fechaInicioParam,
			"$lte": fechaFinParam,
		}
	}
	if tipoDTEParam != "" {
		filter["data.identificacion.tipoDte"] = bson.M{"$gte": tipoDTEParam}
	}
	if condicionOperacionParam != "" {
		filter["data.resumen.condicionOperacion"] = bson.M{"$gte": condicionOperacionParam}
	}

	// Consultar MongoDB con el filtro usando el repositorio
	dtes, err := repositories.GetDTEsByFilter(filter, tipoDTEParam, condicionOperacionParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtes)
}
