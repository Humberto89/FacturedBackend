package controllers

import (
	"Go_Gin/repositories"

	"github.com/gin-gonic/gin"
)

// ObtenerArchivosYDatosJSON es el controlador para obtener archivos y extraer datos.
func ReporteCompra(c *gin.Context) {
	//colecciones a extraer dte
	collections := []string{
		"Factura_de_exportacion",
		"Nota_de_debito",
		"Nota_de_credito",
		"Comprobante_de_credito_fiscal",
	}
	repositories.ReporteCompra(c, collections)
}
