// controllers/dte_controller.go
package controllers

import (
	"net/http"
	"time"

	"Go_Gin/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Controlador para obtener DTEs según parámetros de la URL
func GetDTEs(c *gin.Context) {
	// Obtener parámetros de la URL
	tipoDTE := c.Query("tipoDTE")
	fechaInicioParam := c.Query("fechaInicio")
	fechaFinParam := c.Query("fechaFin")

	// Parsear fechas si están presentes
	var fechaInicio, fechaFin time.Time
	var err error
	if fechaInicioParam != "" {
		fechaInicio, err = time.Parse("2006-01-02", fechaInicioParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha de inicio inválido"})
			return
		}
	}
	if fechaFinParam != "" {
		fechaFin, err = time.Parse("2006-01-02", fechaFinParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha de fin inválido"})
			return
		}
	}

	// Construir filtro para la consulta
	filter := bson.M{}
	if tipoDTE != "" {
		filter["identificacion.tipoDte"] = tipoDTE
	}
	if !fechaInicio.IsZero() {
		filter["identificacion.fecEmi"] = bson.M{"$gte": fechaInicio}
	}
	if !fechaFin.IsZero() {
		// Agregar la condición según sea necesario
		filter["identificacion.fecEmi"] = bson.M{"$lte": fechaFin}
	}

	// Consultar MongoDB con el filtro usando el repositorio
	dtes, err := repositories.GetDTEsByFilter(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtes)
}
