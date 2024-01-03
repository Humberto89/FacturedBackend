package controllers

import (
	"Go_Gin/repositories"
	"log"
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
	filterDTEDate := bson.M{}
	if fechaInicioParam != "" && fechaFinParam != "" {
		filterDTEDate["data.identificacion.fecEmi"] = bson.M{
			"$gte": fechaInicioParam,
			"$lte": fechaFinParam,
		}
	}

	// Obtener todos los documentos de la colección específica
	dtes, err := repositories.GetDTEsByType(filterDTEDate, tipoDTEParam, fechaInicioParam, fechaFinParam, condicionOperacionParam, estadoDTEParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("DTEs encontrados: %v\n", dtes)
	c.JSON(http.StatusOK, dtes)
}
