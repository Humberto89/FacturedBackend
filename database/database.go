package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func CreateDatabase() {
	// Configuración de conexión a MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Comprueba la conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Nombre de la base de datos
	nombreBaseDeDatos := "DTE_Recepcion"

	// Lista de nombres de colecciones
	colecciones := []string{"Archivos",
		"Comprobante_de_credito_fiscal",
		"Comprobante_de_donacion",
		"Comprobante_de_liquidacion",
		"Comprobante_de_retencion",
		"Documento_contable_de_liquidacion",
		"Factura",
		"Factura_de_sujeto_excluido",
		"Factura_de_exportacion",
		"Nota_de_credito",
		"Nota_de_debito",
		"Nota_de_remision"}

	// Crear la base de datos si no existe
	err = client.Database(nombreBaseDeDatos).CreateCollection(context.TODO(), "control")
	if err != nil {
		log.Printf("Error al crear la colección 'control': %v\n", err)
	} else {
		fmt.Printf("Colección 'control' creada con éxito.\n")
	}

	// Crear colecciones si no existen
	for _, nombreColeccion := range colecciones {
		err = client.Database(nombreBaseDeDatos).CreateCollection(context.TODO(), nombreColeccion)
		if err != nil {
			log.Printf("Error al crear la colección %s: %v\n", nombreColeccion, err)
		} else {
			fmt.Printf("Colección %s creada con éxito.\n", nombreColeccion)
		}
	}

	fmt.Println("¡Proceso completado!")
}

func ConnectdbMongo() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Error al conectar a MongoDB: %v", err)
	}

	// Verificar la conexión con una operación simple
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error al verificar la conexión a MongoDB: %v", err)
	}

	fmt.Println("Conexión exitosa a MongoDB")

	collection = client.Database("DTE_Recepcion").Collection("Archivos")
	return collection, nil
}

func ConnectdbPostgre() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "user=postgres password=0000 dbname=DBFormulario sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
