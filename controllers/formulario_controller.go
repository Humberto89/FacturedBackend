// controllers/formulario_controller.go
package controllers

import (
	"Go_Gin/models"

	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func CheckIdentificationExists(c *gin.Context, db *gorm.DB) {
	idType := c.Query("type")
	idValue := c.Query("value")

	if idType == "" || idValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type and value parameters are required"})
		return
	}

	// Convierte el tipo a minúsculas para manejarlo de manera uniforme
	idType = strings.ToLower(idType)

	// Determina la tabla y la columna según el tipo de documento
	tableName, columnName := "", ""
	switch idType {
	case "dui":
		tableName = "formularios"
		columnName = "dui"
	case "nit":
		tableName = "formularios"
		columnName = "nit"
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported type"})
		return
	}

	// Consulta la base de datos para verificar la existencia del documento
	var count int64
	result := db.Table(tableName).Where(fmt.Sprintf("%s = ?", columnName), idValue).Count(&count)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": count > 0})
}

// isIdentificationExists verifica si un documento ya existe en la base de datos.
func isIdentificationExists(db *gorm.DB, idType string, idValue string) bool {
	var count int64
	result := db.Table("formularios").Where(fmt.Sprintf("%s = ?", idType), idValue).Count(&count)
	if result.Error != nil {
		// Manejar el error según tus necesidades (puedes loguearlo o devolver un error)
		fmt.Printf("Error checking existence of %s: %s\n", idType, result.Error.Error())
		return false
	}
	return count > 0
}

func CreateFormulario(c *gin.Context, db *gorm.DB) {
	var formulario models.Formulario
	if err := c.BindJSON(&formulario); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	// Verificar si el DUI ya existe
	if exists := isIdentificationExists(db, "dui", formulario.Dui); exists {
		c.JSON(http.StatusConflict, gin.H{"error": "DUI already exists"})
		return
	}

	// Verificar si el NIT ya existe
	if exists := isIdentificationExists(db, "nit", formulario.NIT); exists {
		c.JSON(http.StatusConflict, gin.H{"error": "NIT already exists"})
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
	fmt.Println("Intento de eliminación para el ID:", id)

	// Convertir el ID de cadena a un tipo numérico (ajusta el tipo según tu modelo)
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.Delete(models.Formulario{}, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Formulario not found"})
			return
		}

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
