package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func main() {

	connectdb()

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5173/*"} //Remplazar por el del front
	r.Use(cors.New(config))

	r.POST("/upload", func(c *gin.Context) {
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
			if !extensionValor(file.Filename) {
				filesError = append(filesError, gin.H{"error": fmt.Sprintf("Tipo de archivo no permitido para %s. Solo se permiten archivos .json y .pdf.", file.Filename)})
				continue
			}

			var isValid bool
			var tipoDteTexto string
			var err error

			// Solo validar la estructura si es un archivo JSON
			if strings.ToLower(filepath.Ext(file.Filename)) == ".json" {
				isValid, tipoDteTexto, err = ValidarEstructuraJSON(file)
				if err != nil {
					filesError = append(filesError, gin.H{
						"filename": file.Filename,
						"error":    fmt.Sprintf("Error al procesar %s: %s", file.Filename),
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

			fileID, err := guardarArchivoMongo(file)
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
				"fileType":   getFileType(file.Filename),
			})
		}

		response := gin.H{
			"filesInfo":  fileInfos,
			"filesError": filesError,
		}

		c.JSON(http.StatusOK, response)
	})

	r.Run(":8081") //Correra en el puerto 8081

}

func guardarArchivoMongo(file *multipart.FileHeader) (string, error) {
	if !extensionValor(file.Filename) {
		return "", fmt.Errorf("tipo de archivo no permitido. Solo se permiten archivos .json y .pdf.")
	}

	fileData, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileData.Close()

	doc := bson.M{
		"filename":   file.Filename,
		"size":       file.Size,
		"uploadDate": time.Now(),
	}

	extension := strings.ToLower(filepath.Ext(file.Filename))

	if extension == ".json" {

		dataBytes, err := ioutil.ReadAll(fileData)
		if err != nil {
			return "", fmt.Errorf("error al leer el archivo JSON: %v", err)
		}

		// Convertir el JSON a la estructura Objeto
		var objeto map[string]interface{}
		if err := json.Unmarshal(dataBytes, &objeto); err != nil {
			return "", fmt.Errorf("error al convertir JSON a objeto: %v", err)
		}

		doc["data"] = objeto

		// Obtener el valor de identificacion/codgeneracion como _id
		if codigoGeneracion, ok := objeto["identificacion"].(map[string]interface{})["codigoGeneracion"].(string); ok {
			doc["_id"] = codigoGeneracion
		} else {
			return "", fmt.Errorf("no se puede obtener el Codigo de Generacion")
		}

		// Obtener el valor de identificacion/tipoDte
		tipoDte, ok := objeto["identificacion"].(map[string]interface{})["tipoDte"].(string)
		if !ok {
			return "", fmt.Errorf("no se puede obtener el TipoDte")
		}

		// Determinar la colección según el tipoDte
		collectionName := "Archivos" // Nombre de la colección por defecto
		switch tipoDte {
		case "01":
			collectionName = "Factura"
		case "03":
			collectionName = "Comprobante_de_credito_fiscal"

		case "04":
			collectionName = "Nota_de_remision"

		case "05":
			collectionName = "Nota_de_credito"

		case "06":
			collectionName = "Nota_de_debito"

		case "07":
			collectionName = "Comprobante_de_retencion"

		case "08":
			collectionName = "Comprobante_de_liquidacion"

		case "09":
			collectionName = "Documento_contable_de_liquidacion"

		case "11":
			collectionName = "Facturas_de_exportacion"

		case "14":
			collectionName = "Factura_de_sujeto_excluido"

		case "15":
			collectionName = "Comprobante_de_donacion"

		}

		// Establecer la colección según el tipoDte
		tipoDteCollection := collection.Database().Collection(collectionName)

		// Insertar en la colección correspondiente
		result, err := tipoDteCollection.InsertOne(context.Background(), doc)
		if err != nil {
			return "", fmt.Errorf("error al insertar en la colección %s: %v", tipoDte, err)
		}

		return fmt.Sprintf("%v", result.InsertedID), nil

	} else if extension == ".pdf" {
		base64String, err := pdfToBase64(fileData)
		if err != nil {
			return "", fmt.Errorf("error al decifrar el PDF: %v", err)

		}

		filenameWithoutExt := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
		doc["_id"] = filenameWithoutExt

		doc["pdfData"] = base64String
	} else {
		return "", fmt.Errorf("tipo de archivo no permitido. Solo se permiten archivos .json y .pdf.")
	}

	result, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result.InsertedID), nil

}

func extensionValor(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json" || ext == ".pdf"
}

func getFileType(filename string) string {
	extension := strings.ToLower(filepath.Ext(filename))
	switch extension {
	case ".json":
		return "JSON"
	case ".pdf":
		return "PDF"
	default:
		return "Desconocido"
	}
}

func pdfToBase64(file multipart.File) (string, error) {
	// Leer el contenido del archivo PDF
	pdfBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Convertir el contenido a base64
	base64String := base64.StdEncoding.EncodeToString(pdfBytes)

	return base64String, nil
}
