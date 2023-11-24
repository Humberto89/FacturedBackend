// controllers/formulario_controller.go
package controllers

import (
	"Go_Gin/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Manejadores CRUD para Formulario

func GetFormulario(c *gin.Context, db *gorm.DB) {
	var formularios []models.Formulario
	if err := db.Preload("Pais").Preload("Departamento").Preload("Municipio").Find(&formularios).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, formularios)
}

func GetFormularioByID(c *gin.Context, db *gorm.DB) {
	var formulario models.Formulario
	id := c.Param("id")
	if err := db.Preload("Pais").Preload("Departamento").Preload("Municipio").First(&formulario, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Formulario not found"})
		return
	}
	c.JSON(200, formulario)
}

func CreateFormulario(c *gin.Context, db *gorm.DB) {
	var formulario models.Formulario
	if err := c.BindJSON(&formulario); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	if err := db.Create(&formulario).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, formulario)
}

func UpdateFormulario(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var formulario models.Formulario
	if err := db.First(&formulario, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Formulario not found"})
		return
	}
	if err := c.BindJSON(&formulario); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	if err := db.Save(&formulario).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, formulario)
}

func DeleteFormulario(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(models.Formulario{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Formulario deleted successfully"})
}

// En tus controladores
func GetPaises(c *gin.Context, db *gorm.DB) {
	var paises []models.Pais

	if err := db.Find(&paises).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, paises)
}

func GetDepartamentos(c *gin.Context, db *gorm.DB) {
	paisID := c.Param("paisID")

	var departamentos []models.Departamento
	if err := db.Where("pais_id = ?", paisID).Find(&departamentos).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, departamentos)
}

func GetMunicipios(c *gin.Context, db *gorm.DB) {
	departamentoID := c.Param("departamentoID")

	var municipios []models.Municipio
	if err := db.Where("departamento_id = ?", departamentoID).Find(&municipios).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, municipios)
}
