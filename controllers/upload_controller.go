package controllers

import (
	"Go_Gin/database"
	"Go_Gin/repositories"
	"Go_Gin/services"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func FilesUpload(c *gin.Context) {
	// Obtener el token del encabezado
	token := c.GetHeader("Authorization")

	// Validar el token
	if err := ValidateToken(token); err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Extraer el campo groupsid del token
	IdentifierEmp, err := services.ExtraerIdEmpr(token)
	if err != nil {
		// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	client, err := database.ConnectdbMongo()
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("DTE_Recepcion").Collection("Archivos")

	form, _ := c.MultipartForm()
	files := form.File["files"]

	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ningún archivo seleccionado"})
		return
	}

	var fileInfos []gin.H  // Almacena la información de cada archivo
	var filesError []gin.H // Almacena cada uno de los errores

	for _, file := range files {
		log.Printf("Proccessing file: %s\n", file.Filename)
		if !services.ExtensionValor(file.Filename) {
			filesError = append(filesError, gin.H{"error": fmt.Sprintf("Tipo de archivo no permitido para %s. Solo se permiten archivos .json y .pdf.", file.Filename)})
			continue
		}

		var isValid bool
		var tipoDteTexto string
		var err error

		// Solo validar la estructura si es un archivo JSON
		if strings.ToLower(filepath.Ext(file.Filename)) == ".json" {
			isValid, tipoDteTexto, err = services.ValidarEstructuraJSON(file)
			if err != nil {
				filesError = append(filesError, gin.H{
					"filename": file.Filename,
					"error":    fmt.Sprintf("error al procesar %s: %s ", file.Filename, ""),
				})
				continue
			}

			if !isValid {
				filesError = append(filesError, gin.H{
					"filename": file.Filename,
					"error":    fmt.Sprintf("El archivo '%s' no cumple con el esquema JSON", file.Filename),
				})
				continue
			}

		}

		fileID, err := repositories.GuardarArchivoMongo(file, collection, IdentifierEmp)
		if err != nil {
			filesError = append(filesError, gin.H{
				"filename": file.Filename,
				"error":    fmt.Sprintf("Error al procesar %s: %s", file.Filename, "Ya se encuentra en la base de datos"),
			})
			continue
		}

		fileInfos = append(fileInfos, gin.H{
			"filename":   file.Filename,
			"message":    "Archivo " + file.Filename + " subido correctamente",
			"tipoDte":    tipoDteTexto,
			"mongoDB_ID": fileID,
			"fileType":   services.GetFileType(file.Filename),
		})
	}

	response := gin.H{
		"filesInfo":  fileInfos,
		"filesError": filesError,
	}

	c.JSON(http.StatusOK, response)
}
