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

	// Construir filtros para la consulta
	filter := bson.M{}
	filterDTEType := bson.M{}
	filterDteDate := bson.M{}
	//definiendo valores
	if fechaInicioParam != "" && fechaFinParam != "" {
		// Usar notación de puntos para filtrar por fecha dentro del objeto identificacion
		filterDteDate["data.identificacion.fecEmi"] = bson.M{
			"$gte": fechaInicioParam,
			"$lte": fechaFinParam,
		}
	}
	if tipoDTEParam != "" {
		filter["data.identificacion.tipoDte"] = bson.M{"$gte": tipoDTEParam}
	}
	//filtro para tipo de DTE
	if condicionOperacionParam != "" {
		//accediendo al tipo de pago
		filterDTEType["data.resumen.condiconOperacion"] = bson.M{"$gte": condicionOperacionParam}

	}

	// Consultar MongoDB con el filtro usando el repositorio
	dtes, err := repositories.GetDTEsByType(filterDteDate, tipoDTEParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtes)

}
