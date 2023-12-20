// controllers/dte_controller.go
package controllers

import (
	"Go_Gin/repositories"
	"fmt"
	"net/http"

	"strconv"

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

	// Convertir condicionOperacionParam a entero
	condicionOperacion, err := strconv.Atoi(condicionOperacionParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al convertir condicionOperacionParam a entero: %v", err)})
		return
	}
	//validar el valor de la condicion de la operacion
	if condicionOperacion < 1 || condicionOperacion > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El valor de la condición de operación no es válido"})
		return
	}
	// Construir filtros para la consulta
	//filterDTEOp := bson.M{}
	filterDTEType := bson.M{}
	filterDTEOp := bson.M{}
	filterDTEDate := bson.M{}
	//definiendo valores
	if fechaInicioParam != "" && fechaFinParam != "" {
		// Usar notación de puntos para filtrar por fecha dentro del objeto identificacion
		filterDTEDate["data.identificacion.fecEmi"] = bson.M{
			"$gte": fechaInicioParam,
			"$lte": fechaFinParam,
		}
	}
	if tipoDTEParam != "" {
		filterDTEType["data.identificacion.tipoDte"] = bson.M{"$gte": tipoDTEParam}
	}
	//filtro para tipo de DTE
	if condicionOperacion != 0 {
		// Accediendo al tipo de pago
		filterDTEOp["resumen.condicionOperacion"] = bson.M{"resumen.condicionOperacion": condicionOperacion}
	}

	// Consultar MongoDB con el filtro usando el repositorio
	dtes, err := repositories.GetDTEsByType(filterDTEDate, tipoDTEParam, condicionOperacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtes)

}
