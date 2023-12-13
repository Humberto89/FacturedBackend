package repositories

import (
	"Go_Gin/database"
	"Go_Gin/models"
	"context"
	"fmt"
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
var operationsCM = map[string]string{
	"1": "Contado",
	"2": "A credito",
	"3": "Otro",
}

func GetDTEsByFilter(filter bson.M, tipoDTE string, operacion string) ([]models.Documento, error) {
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

	collection := client.Database("DTE_Recepcion").Collection(dteColeccion)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Agregar condición de operación al filtro si está presente
	if operacion != "" {
		filter["data.resumen.condicionOperacion"] = bson.M{"$gte": operacion}
	}

	// Consulta para el tipoDTE (y la condición de operación si está presente)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("Error al realizar la búsqueda: %v", err)
	}
	defer cursor.Close(ctx)

	var resultados []models.Documento
	if err := cursor.All(ctx, &resultados); err != nil {
		return nil, fmt.Errorf("Error al decodificar los resultados: %v", err)
	}

	return resultados, nil
}
