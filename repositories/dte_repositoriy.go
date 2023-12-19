package repositories

import (
	"Go_Gin/services"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GuardarArchivoMongo(file *multipart.FileHeader, collection *mongo.Collection) (string, error) {
	if !services.ExtensionValor(file.Filename) {
		return "", fmt.Errorf("tipo de archivo no permitido, solo se permiten archivos .json y .pdf")
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
		collectionName := "ArchivosPDF" // Nombre de la colección por defecto
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
			collectionName = "Factura_de_exportacion"

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
		base64String, err := services.PdfToBase64(fileData)
		if err != nil {
			return "", fmt.Errorf("error al decifrar el PDF: %v", err)

		}

		filenameWithoutExt := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
		doc["_id"] = filenameWithoutExt

		doc["pdfData"] = base64String
	} else {
		return "", fmt.Errorf("tipo de archivo no permitido, solo se permiten archivos .json y .pdf")
	}

	result, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result.InsertedID), nil
}
