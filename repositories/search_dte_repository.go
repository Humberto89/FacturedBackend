// repositories/dte_repository.go

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

var collectionMap = map[string]string{
	"01": "Factura",
	"04": "Nota_de_remision",
	"15": "Comprobante_de_donacion",
}

func GetDTEsByFilter(filter bson.M, codigo string) ([]models.Documento, error) {
	// Obtener la colección y realizar la búsqueda
	client, errr := database.ConnectdbMongo()
	if errr != nil {
		log.Fatal(errr)
	}

	// Verificar si el código existe en el mapa
	nombreColeccion, ok := collectionMap[codigo]
	if !ok {
		return nil, fmt.Errorf("Código de colección no válido")
	}

	collection := client.Database("DTE_Recepcion").Collection(nombreColeccion)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
