package controllers

import (
	"Go_Gin/repositories"

	"github.com/gin-gonic/gin"
)

// ObtenerArchivosYDatosJSON es el controlador para obtener archivos y extraer datos.
func ReporteCompra(c *gin.Context) {
	repositories.ReporteCompra(c)
}
