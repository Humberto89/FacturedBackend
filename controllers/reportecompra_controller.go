package controllers

import (
	"Go_Gin/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ObtenerArchivosYDatosJSON es el controlador para obtener archivos y extraer datos.
func ReporteAnexo(c *gin.Context) {
	// Obtener el token del encabezado
	token := c.GetHeader("Authorization")

	// Validar el token
	if err := ValidateToken(token); err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

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
