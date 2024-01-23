package controllers

import (
	"Go_Gin/database"
	"Go_Gin/repositories"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"
)

// Controlador para obtener DTEs según parámetros de la URL
func GetDTEs(c *gin.Context, db *gorm.DB) {
	// Obtener parámetros de la URL
	tipoDTEParam := c.Query("tipoDTE")
	estadoDTEParam := c.Query("estadoSeguimiento")
	fechaInicioParam := c.Query("fechaInicio")
	fechaFinParam := c.Query("fechaFin")
	condicionOperacionParam := c.Query("condicionOperacion")
	identifierEmp := c.GetHeader("IdentifierEmp")

	// Construir filtros para la consulta
	filterDTEDate := bson.M{}
	if fechaInicioParam != "" && fechaFinParam != "" {
		filterDTEDate["data.identificacion.fecEmi"] = bson.M{
			"$gte": fechaInicioParam,
			"$lte": fechaFinParam,
		}
	}

	// Obtener todos los documentos de la colección específica
	dtes, err := repositories.GetDTEsByType(filterDTEDate, tipoDTEParam, fechaInicioParam, fechaFinParam, condicionOperacionParam, estadoDTEParam, identifierEmp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("DTEs encontrados: %v\n", dtes)
	c.JSON(http.StatusOK, dtes)
}

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

// Función para manejar la solicitud y actualizar el estadoSeguimiento
func HandleActualizarEstadoSeguimiento(c *gin.Context) {
	// Parsear los parámetros de la URL
	id := c.Query("id")
	opcion := c.Query("opcion")
	tipoDte := c.Query("tipoDte")

	// Validar que se proporcionaron los parámetros necesarios
	if id == "" || opcion == "" || tipoDte == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requieren los parámetros 'id', 'opcion' y 'tipoDte'"})
		return
	}

	// Realizar el mapeo de datos basado en el valor de tipoDte
	collection, ok := collectionMap[tipoDte]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No se encontró mapeo para tipoDte: %s", tipoDte)})
		return
	}

	// Convertir opcion a entero
	nuevoEstado, err := strconv.Atoi(opcion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La opción debe ser un número entero"})
		return
	}

	// Obtener la colección y realizar la búsqueda
	client, err := database.ConnectdbMongo()
	if err != nil {
		log.Printf("Error al conectar a MongoDB: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Seleccionar la colección
	mongoCollection := client.Database("DTE_Recepcion").Collection(collection)

	// Crear el filtro para encontrar el documento por ID
	filter := bson.D{{Key: "_id", Value: id}}

	// Crear la actualización para modificar el estadoSeguimiento
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "estadoSeguimiento", Value: nuevoEstado}}}}

	// Actualizar el documento en la base de datos
	result, err := mongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error al actualizar el estadoSeguimiento: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	// Verificar si se realizó la actualización
	if result.ModifiedCount == 0 {
		log.Printf("No se encontró el documento con ID %s en la colección %s", id, tipoDte)
		c.JSON(http.StatusNotFound, gin.H{"error": "Documento no encontrado"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "EstadoSeguimiento actualizado correctamente"})
}
