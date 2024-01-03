package controllers

import (
	"Go_Gin/repositories"

	"github.com/gin-gonic/gin"
)

// ObtenerArchivosYDatosJSON es el controlador para obtener archivos y extraer datos.
func ReporteAnexo(c *gin.Context) {
	//colecciones a extraer dte
	collections := []string{
		"Nota_de_remision",
		"Comprobante_de_donacion",
		"Factura_de_sujeto_excluido",
		"Comprobante_de_liquidacion",
		"Comprobante_de_retencion",
		"Documento_contable_de_liquidacion",
		"Factura",
	}

	repositories.ReporteAnexo(c, collections)
}
