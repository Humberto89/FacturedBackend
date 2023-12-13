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

func GetDTEsByFilter(filter bson.M, codigo string, operacion string) ([]models.Documento, error) {
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

	//tipo de operacion
	operacionColeccion, ok := operationsCM[operacion]
	if !ok {
		return nil, fmt.Errorf("Codigo de coleccion no valido")
	}
	opCollection := client.Database("DTE_Recepcion").Collection(operacionColeccion)
	ctx, cancelar := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelar()

	cursorOp, err := opCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("Error a realizar la busqueda %v", err)
	}
	defer cursorOp.Close(ctx)

	if err := cursor.All(ctx, &resultados); err != nil {
		return nil, fmt.Errorf("Error al decodificar los resultados: %v", err)
	}

	return resultados, nil
}
