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
	estadoDTEParam := c.Query("estadoSeguimiento")
	fechaInicioParam := c.Query("fechaInicio")
	fechaFinParam := c.Query("fechaFin")
	condicionOperacionParam := c.Query("condicionOperacion")
	// Construir filtros para la consulta
	filterDTEType := bson.M{}
	filterDTEOp := bson.M{}
	filterDTEDate := bson.M{}
	filterStatusDTE := bson.M{}
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
	if condicionOperacionParam != "" {
		// Accediendo al tipo de pago
		filterDTEOp["data.resumen.condicionOperacion"] = bson.M{"$gte": condicionOperacionParam}
	}
	if estadoDTEParam != "" {
		filterStatusDTE["estadoSeguimiento"] = bson.M{"$gte": estadoDTEParam}
	}
	// Consultar MongoDB con el filtro usando el repositorio
	dtes, err := repositories.GetDTEsByType(filterDTEDate, tipoDTEParam, condicionOperacionParam, estadoDTEParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtes)

}
