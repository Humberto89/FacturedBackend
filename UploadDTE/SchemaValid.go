package main

import (
	"Go_Gin/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

var ident models.Ident

func ValidarEstructuraJSON(file *multipart.FileHeader) (bool, string, error) {
	// Abrir el archivo original recibido
	originalFile, err := file.Open()
	if err != nil {
		return false, "", fmt.Errorf("Error al abrir el archivo original: %v", err)
	}
	defer originalFile.Close()

	// Leer el contenido del archivo JSON
	dataBytes, err := ioutil.ReadAll(originalFile)
	if err != nil {
		return false, "", fmt.Errorf("Error al leer el contenido del archivo JSON: %v", err)
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
			return false, "", fmt.Errorf("Error al obtener la ruta absoluta para el esquema: %v", err)
		}

		// Cargar el esquema desde un archivo JSON
		schemaLoader := gojsonschema.NewReferenceLoader("file://" + filepath.ToSlash(schemaPath))

		// Validar los datos contra el esquema
		result, err := gojsonschema.Validate(schemaLoader, dataLoader)
		if err != nil {
			return false, "", fmt.Errorf("Error al validar: %v", err)
		}

		if result.Valid() {
			// Decodificar el JSON para obtener el campo TipoDte
			// var data Documento
			if err := json.Unmarshal(dataBytes, &ident); err != nil {
				return false, "", fmt.Errorf("Error al decodificar el JSON: %v", err)
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
