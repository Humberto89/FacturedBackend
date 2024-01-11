package repositories

import (
	"Go_Gin/database"
	"Go_Gin/models"
	"context"
	"fmt"
	"log"
	"strconv"
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

// Filtrar por tipo de DTE, estadoDTE y condicionOperacion
func GetDTEsByType(filterDTEDate bson.M, tipoDTE string, fechaInicio string, fechaFin string, condicionOperacion string, estadoDTE string, identifierEmp string) ([]models.Documento, error) {
	// Obtener la colección y realizar la búsqueda
	client, err := database.ConnectdbMongo()
	if err != nil {
		log.Fatal(err)
	}

	// Obtener el nombre de la colección basado en el tipoDTE
	dteColeccion, ok := collectionMap[tipoDTE]
	if !ok {
		return nil, fmt.Errorf("TipoDTE no válido")
	}

	// Obtener la colección específica
	collectionType := client.Database("DTE_Recepcion").Collection(dteColeccion)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Aplicar filtro de fecha si se proporcionan fechaInicio y fechaFin
	if fechaInicio != "" && fechaFin != "" {
		filterDTEDate["data.identificacion.fecEmi"] = bson.M{
			"$gte": fechaInicio,
			"$lte": fechaFin,
		}
	}

	// Aplicar filtro de EstadoSeguimiento si se proporciona
	if estadoDTE != "" {
		// Convierte el valor de estadoDTE a int
		estadoDTEInt, err := strconv.Atoi(estadoDTE)
		if err != nil {
			log.Printf("Error al convertir estadoDTE a int: %v\n", err)
			return nil, fmt.Errorf("error al convertir estadoDTE a int: %v", err)
		}
		filterDTEDate["estadoSeguimiento"] = estadoDTEInt
	}

	// Aplicar filtro de condicionOperacion si se proporciona
	if condicionOperacion != "" {
		// Convierte el valor de condicionOperacion a int
		condicionOperacionInt, err := strconv.Atoi(condicionOperacion)
		if err != nil {
			log.Printf("Error al convertir condicionOperacion a int: %v\n", err)
			return nil, fmt.Errorf("error al convertir condicionOperacion a int: %v", err)
		}
		filterDTEDate["data.resumen.condicionOperacion"] = condicionOperacionInt
	}

	log.Printf("Filtro aplicado: %v\n", filterDTEDate)

	filterDTEDate["EmpID"] = identifierEmp

	// Consultar MongoDB con el filtro
	cursorType, err := collectionType.Find(ctx, filterDTEDate)
	if err != nil {
		log.Printf("Error al realizar la busqueda en MongoDB: %v\n", err)
		return nil, fmt.Errorf("error al realizar la búsqueda: %v", err)
	}
	defer cursorType.Close(ctx)

	// Decodificar resultados para tipo de DTE
	var resultados []models.Documento
	if err := cursorType.All(ctx, &resultados); err != nil {
		return nil, fmt.Errorf("error al decodificar los resultados: %v", err)
	}

	log.Printf("resultados encontrados: %v\n", resultados)
	return resultados, nil
}
