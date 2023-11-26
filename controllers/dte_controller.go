package controllers

import (
	"Go_Gin/models"
	"context"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func ObtenerTipoDTE() string {
	tipoDTE := []string{"01", "03", "04", "05", "06", "07", "08", "09", "11", "14", "15"}
	index := rand.Intn(len(tipoDTE))
	return tipoDTE[index]
}

// getDTEs is a handler function to get all DTE documents from MongoDB
func GetDTEs(c *gin.Context) {
	var dtes []models.Ident

	// Consult MongoDB to get all documents
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error al mostrar los documentos": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	// Iterate over documents and add them to the slice
	for cursor.Next(context.Background()) {
		var ident models.Ident
		if err := cursor.Decode(&ident); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		dtes = append(dtes, ident)
	}
	ObtenerTipoDTE() // This line doesn't seem to have any effect, consider removing or updating it

	c.JSON(http.StatusOK, dtes)
}
