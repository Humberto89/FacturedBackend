// controllers/formulario_controller.go
package controllers

import (
	"Go_Gin/models"
	"Go_Gin/services"

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
	// Obtener el token del encabezado
	token := c.GetHeader("Authorization")

	// Validar el token
	if err := ValidateToken(token); err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var formularios []models.Formulario

	// Extraer el campo groupsid del token
	identifierEmp, err := services.ExtraerIdEmpr(token)
	if err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Aplicar el filtro si se proporciona
	if identifierEmp != "" {
		if err := db.Preload("Pais").Preload("Departamento").Preload("Municipio").
			Where("emp_id = ?", identifierEmp).Find(&formularios).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := db.Preload("Pais").Preload("Departamento").Preload("Municipio").Find(&formularios).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(200, formularios)
}

func GetFormularioByID(c *gin.Context, db *gorm.DB) {
	// Obtener el token del encabezado
	token := c.GetHeader("Authorization")

	// Validar el token
	if err := ValidateToken(token); err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var formulario models.Formulario
	id := c.Param("id")

	// Extraer el campo groupsid del token
	identifierEmp, err := services.ExtraerIdEmpr(token)
	if err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Aplicar el filtro si se proporciona
	query := db.Preload("Pais").Preload("Departamento").Preload("Municipio")
	if identifierEmp != "" {
		query = query.Where("emp_id = ?", identifierEmp)
	}

	if err := query.First(&formulario, id).Error; err != nil {
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
		// Si el tipo es DUI, y el valor es nulo o cadena vacía, no realiza la verificación y devuelve false
		if idValue == "null" || idValue == "" {
			c.JSON(http.StatusOK, gin.H{"exists": false})
			return
		}
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
	// Si el valor es nulo o vacío, devuelve false
	if idValue == "" {
		return false
	}

	var count int64
	result := db.Table("formularios").Where(fmt.Sprintf("%s = ?", idType), idValue).Count(&count)
	if result.Error != nil {
		// Manejar el error según tus necesidades
		fmt.Printf("Error checking existence of %s: %s\n", idType, result.Error.Error())
		return false
	}
	return count > 0
}

func CreateFormulario(c *gin.Context, db *gorm.DB) {
	// Obtener el token del encabezado
	token := c.GetHeader("Authorization")

	// Validar el token
	if err := ValidateToken(token); err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

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

	// Después de la verificación de existencia para DUI y NIT
	fmt.Printf("Attempting to create Formulario with DUI: %s, NIT: %s\n", formulario.Dui, formulario.NIT)

	// Extraer el campo groupsid del token
	identifierEmp, err := services.ExtraerIdEmpr(token)
	if err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	formulario.EmpID = identifierEmp

	if err := db.Create(&formulario).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, formulario)
}

func UpdateFormulario(c *gin.Context, db *gorm.DB) {
	// Obtener el token del encabezado
	token := c.GetHeader("Authorization")

	// Validar el token
	if err := ValidateToken(token); err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var formulario models.Formulario

	// Extraer el campo groupsid del token
	identifierEmp, err := services.ExtraerIdEmpr(token)
	if err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Aplicar el filtro si se proporciona
	query := db
	if identifierEmp != "" {
		query = query.Where("emp_id = ?", identifierEmp)
	}

	if err := query.First(&formulario, id).Error; err != nil {
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
	// Obtener el token del encabezado
	token := c.GetHeader("Authorization")

	// Validar el token
	if err := ValidateToken(token); err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	fmt.Println("Intento de eliminación para el ID:", id)

	// Convertir el ID de cadena a un tipo numérico (ajustar el tipo según tu modelo)
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var formulario models.Formulario
	// Extraer el campo groupsid del token
	identifierEmp, err := services.ExtraerIdEmpr(token)
	if err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el formulario existe y si el emp_id coincide
	if err := db.Where("id = ? AND emp_id = ?", ID, identifierEmp).First(&formulario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Formulario not found"})
			return
		}

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Eliminar el formulario
	if err := db.Delete(&formulario).Error; err != nil {
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
