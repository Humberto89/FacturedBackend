// repositories/dte_repository.go

package repositories

import (
	"Go_Gin/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func SetDTECollection(coll *mongo.Collection) {
	collection = coll
}

func GetDTEsByFilter(filter bson.M) ([]models.Ident, error) {
	var dtes []models.Ident

	// Realizar la consulta a MongoDB usando el filtro
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterar sobre los documentos y agregarlos a la lista
	for cursor.Next(context.Background()) {
		var ident models.Ident
		if err := cursor.Decode(&ident); err != nil {
			return nil, err
		}
		dtes = append(dtes, ident)
	}

	return dtes, nil
}
