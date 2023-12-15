package controllers

import (
	"Go_Gin/models"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetMunicipioByID(c *gin.Context, db *gorm.DB) {
	municipioID := c.Param("id")

	var municipio models.Municipio
	if err := db.First(&municipio, municipioID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Municipio no encontrado"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, municipio)
}

func GetDepartamentoByID(c *gin.Context, db *gorm.DB) {
	departamentoID := c.Param("id")

	var departamento models.Departamento
	if err := db.First(&departamento, departamentoID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Departamento no encontrado"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, departamento)
}

func GetPaisByID(c *gin.Context, db *gorm.DB) {
	paisID := c.Param("id")

	var pais models.Pais
	if err := db.First(&pais, paisID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Pais no encontrado"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pais)
}
