// repositories/dte_repository.go

package repositories

import (
	"Go_Gin/models"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func SetDTECollection(coll *mongo.Collection) {
	collection = coll
}

func GetDTEsByFilter(filter bson.M) ([]models.Ident, error) {
	if collection == nil {
		return nil, errors.New("La colección no ha sido inicializada. Asegúrate de llamar a ConnectdbMongo antes de utilizarla")
	}

	var dtes []models.Ident

	// Realizar la consulta a MongoDB usando el filtro
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta en MongoDB: %v", err)
	}
	defer cursor.Close(context.Background())

	// Iterar sobre los documentos y agregarlos a la lista
	for cursor.Next(context.Background()) {
		var ident models.Ident
		if err := cursor.Decode(&ident); err != nil {
			return nil, fmt.Errorf("Error al decodificar el documento: %v", err)
		}
		dtes = append(dtes, ident)
	}

	return dtes, nil
}
