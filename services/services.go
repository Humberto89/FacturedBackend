package services

import (
	"Go_Gin/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ident models.Ident

func PdfToBase64(file multipart.File) (string, error) {
	// Leer el contenido del archivo PDF
	pdfBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Convertir el contenido a base64
	base64String := base64.StdEncoding.EncodeToString(pdfBytes)

	return base64String, nil
}

func ExtensionValor(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json" || ext == ".pdf"
}

func GetFileType(filename string) string {
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

func getTxtTipoDte(tipoDte string) string {
	switch tipoDte {
	case "01":
		return "Factura"
	case "03":
		return "Comprobante de crédito fiscal"
	case "04":
		return "Nota de remisión"
	case "05":
		return "Nota de crédito"
	case "06":
		return "Nota de débito"
	case "07":
		return "Comprobante de retención"
	case "08":
		return "Comprobante de liquidación"
	case "09":
		return "Documento contable de liquidación"
	case "11":
		return "Facturas de exportación"
	case "14":
		return "Factura de sujeto excluido"
	case "15":
		return "Comprobante de donación"
	default:
		return "Desconocido"
	}
}

func ValidarEstructuraJSON(file *multipart.FileHeader) (bool, string, error) {
	// Abrir el archivo original recibido
	originalFile, err := file.Open()
	if err != nil {
		return false, "", fmt.Errorf("error al abrir el archivo original: %v", err)
	}
	defer originalFile.Close()

	// Leer el contenido del archivo JSON
	dataBytes, err := ioutil.ReadAll(originalFile)
	if err != nil {
		return false, "", fmt.Errorf("error al leer el contenido del archivo JSON: %v", err)
	}

	// Cargar los datos que se desea validar desde el contenido del archivo JSON
	dataLoader := gojsonschema.NewStringLoader(string(dataBytes))

	// Lista de rutas a los esquemas
	schemaPaths := []string{"schema/fe-ccf-v3.json",
		"schema/fe-cd-v1.json",
		"schema/fe-cl-v1.json",
		"schema/fe-cr-v1.json",
		"schema/fe-dcl-v1.json",
		"schema/fe-fc-v1.json",
		"schema/fe-fex-v1.json",
		"schema/fe-fse-v1.json",
		"schema/fe-nc-v3.json",
		"schema/fe-nd-v3.json",
		"schema/fe-nr-v3.json",
	}

	// Validar los datos contra cada esquema
	var isValid bool
	var tipoDteTexto string
	for _, schemaPath := range schemaPaths {
		schemaPath, err := filepath.Abs(schemaPath)
		if err != nil {
			return false, "", fmt.Errorf("error al obtener la ruta absoluta para el esquema: %v", err)
		}

		// Cargar el esquema desde un archivo JSON
		schemaLoader := gojsonschema.NewReferenceLoader("file://" + filepath.ToSlash(schemaPath))

		// Validar los datos contra el esquema
		result, err := gojsonschema.Validate(schemaLoader, dataLoader)
		if err != nil {
			return false, "", fmt.Errorf("error al validar: %v", err)
		}

		if result.Valid() {
			// Decodificar el JSON para obtener el campo TipoDte
			// var data Documento
			if err := json.Unmarshal(dataBytes, &ident); err != nil {
				return false, "", fmt.Errorf("error al decodificar el JSON: %v", err)
			}

			// Obtener el tipoDte del documento
			tipoDte := ident.Identificacion.TipoDte
			// Obtener el texto asociado al tipoDte
			tipoDteTexto = getTxtTipoDte(tipoDte)

			isValid = true
			break
		}
	}

	return isValid, tipoDteTexto, nil
}

// Pdf representa la estructura del documento en MongoDB
type Pdf struct {
	ID       string `json:"_id" bson:"_id"`
	PdfData  string `json:"pdfData" bson:"pdfData"`
	Filename string `json:"filename" bson:"filename"`
	// Agrega más campos según sea necesario
}

// PdfDataGet busca un documento por _id en MongoDB
func PdfDataGet(c *gin.Context) {
	// Inicializar la conexión a MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar a MongoDB"})
		return
	}

	// Conectar al cliente
	err = client.Connect(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar a MongoDB"})
		return
	}
	defer client.Disconnect(context.Background())

	// Seleccionar la base de datos y la colección
	db := client.Database("DTE_Recepcion")
	collection := db.Collection("Archivos")

	// Obtener el parámetro _id de la solicitud
	_id := c.Param("id")

	// Verificar si el parámetro _id está presente
	if _id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetro _id faltante"})
		return
	}

	// Crear un contexto con timeout para la operación de búsqueda en MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Crear un filtro para buscar el documento por _id
	filter := bson.M{"_id": _id}

	// Realizar la consulta en la base de datos
	var resultado Pdf
	err = collection.FindOne(ctx, filter).Decode(&resultado)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Documento no encontrado"})
		return
	}

	// Enviar la respuesta en formato JSON
	c.JSON(http.StatusOK, resultado)
}

func ExtraerIdEmpr(tokenString string) (string, error) {
	// Verificar y quitar el prefijo "Bearer "
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Procesar el token como antes
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("Token no válido")
	}

	decodedPayload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("Error al decodificar el payload: %v", err)
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(decodedPayload, &claims); err != nil {
		return "", fmt.Errorf("Error al decodificar el payload JSON: %v", err)
	}

	groupID, ok := claims["groupsid"].(string)
	if !ok {
		return "", fmt.Errorf("Campo 'groupsid' no encontrado en los claims")
	}

	return groupID, nil
}
