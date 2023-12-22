package repositories

import (
	"Go_Gin/database"
	"Go_Gin/models"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// mapeo de las colecciones
var collectionMap = map[string]string{
	"01": "Factura",
	"03": "Comprobante_de_credito_fiscal",
	"04": "Nota_de_remision",
	"05": "Nota_de_credito",
	"06": "Nota_de_debito",
	"07": "Comprobante_de_retencion",
	"08": "Comprobante_de_liquidacion",
	"09": "Documento_contable_de_liquidacion",
	"11": "Factura_de_exportacion",
	"14": "Factura_de_sujeto_excluido",
	"15": "Comprobante_de_donacion",
}

// mapeo de la coleccion de las operaciones
var operationMap = map[int]string{
	1: "1",
	2: "2",
	3: "3",
}

// mapeo de la coleccion de las operaciones
var statusMap = map[int]string{
	1: "1",
	2: "2",
	3: "3",
}

// filtrar por tipo de DTE
func GetDTEsByType(filterDTEDate bson.M, tipoDTE string, condicionOperacion bson.M, estadoDTE bson.M) ([]models.Documento, error) {
	// Obtener la colección y realizar la búsqueda
	client, err := database.ConnectdbMongo()
	if err != nil {
		log.Fatal(err)
	}

	// Verificar si el tipoDTE existe en el mapa
	dteColeccion, ok := collectionMap[tipoDTE]
	if !ok {
		return nil, fmt.Errorf("TipoDTE no válido")
	}

	//tipo de DTE
	collectionType := client.Database("DTE_Recepcion").Collection(dteColeccion)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Agregar condición de operación al filtro si está presente
	// Consulta para el tipoDTE (y la condición de operación si está presente)
	cursorType, err := collectionType.Find(ctx, filterDTEDate)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la búsqueda: %v", err)
	}
	defer cursorType.Close(ctx)
	//consulta de condicion de operacion
	cursorOp, err := collectionType.Find(ctx, condicionOperacion)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la busqueda: %v", err)
	}
	defer cursorOp.Close(ctx)
	//consulta de condicion de operacion
	cursorSt, err := collectionType.Find(ctx, estadoDTE)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la busqueda: %v", err)
	}
	defer cursorSt.Close(ctx)
	//decodificacion de resultados para tipo de DTE
	var resultados []models.Documento
	if err := cursorType.All(ctx, &resultados); err != nil {
		return nil, fmt.Errorf("error al decodificar los resultados: %v", err)
	}
	// Decodificación de resultados para condición de operación y agregado al slice
	var resultadosOp []models.Documento
	if err := cursorOp.All(ctx, &resultadosOp); err != nil {
		return nil, fmt.Errorf("error al decodificar los resultados para condición de operación: %v", err)
	}
	// Decodificación de resultados para condición de operación y agregado al slice
	var resultadosSt []models.Documento
	if err := cursorSt.All(ctx, &resultadosSt); err != nil {
		return nil, fmt.Errorf("error al decodificar los resultados para condición de operación: %v", err)
	}
	//combinando salidas
	resultados = append(resultados, resultadosOp...)
	resultados = append(resultados, resultadosSt...)
	log.Printf("resultados encontrados: %v\n", resultados)
	return resultados, nil
}

// Fase 2
// Convertir el contenido base64 a bytes
func base64ToBytes(base64String string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar base64: %v", err)
	}
	return decodedBytes, nil
}

// Escribir los bytes en un archivo PDF
func bytesToPDF(bytes []byte, filename string) error {
	return ioutil.WriteFile(filename, bytes, 0644)
}

// Función para obtener datos del PDF por _id
func GetPDFDataByID(id string) error {
	// Obtener la colección directamente dentro de la función
	client, err := database.ConnectdbMongo()
	if err != nil {
		return fmt.Errorf("error al conectar a MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())
	// Obtener la colección "Archivos"
	collection := client.Database("DTE_Recepcion").Collection("Archivos")

	// Buscar el documento por _id
	var documento bson.M
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&documento)
	if err != nil {
		return fmt.Errorf("error al buscar el documento: %v", err)
	}

	// Acceder directamente al campo pdfData y realizar la decodificación
	pdfData, ok := documento["pdfData"].(string)
	if !ok {
		return fmt.Errorf("campo pdfData no encontrado o no es una cadena")
	}

	pdfBytes, err := base64.StdEncoding.DecodeString(pdfData)
	if err != nil {
		return fmt.Errorf("error al decodificar base64: %v", err)
	}

	// Escribir los bytes en un nuevo archivo PDF
	err = bytesToPDF(pdfBytes, "nuevo_archivo.pdf")
	if err != nil {
		return fmt.Errorf("error al escribir el archivo PDF: %v", err)
	}

	fmt.Println("Conversión exitosa. Compara el nuevo_archivo.pdf con el original.")
	return nil
}

// Función para obtener datos del PDF por código de generación
func GetPDFData(id string) {
	// Llamada a la función GetPDFDataByID con el _id proporcionado
	err := GetPDFDataByID(id)
	if err != nil {
		fmt.Println(err)
	}
}
