package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectdb() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Verificar la conexión con una operación simple
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error al verificar la conexión a MongoDB:", err)
	} else {
		fmt.Println("Conexión exitosa a MongoDB")
	}

	collection = client.Database("ArchivosPrueba").Collection("Archivos")

}
