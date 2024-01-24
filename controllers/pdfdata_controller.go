package controllers

import (
	"Go_Gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PdfDataGet(c *gin.Context) {

	// Obtener el token del encabezado
	token := c.GetHeader("Authorization")

	// Validar el token
	if err := ValidateToken(token); err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	services.PdfDataGet(c)

}
