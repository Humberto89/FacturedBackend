package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

// DTE representa la estructura de un Documento Tributario Electrónico
type DTE struct {
	ID                   string            `json:"id,omitempty"`
	Identificador        Identificador     `json:"identificador"`
	DocumentoRelacionado string            `json:"documentoRelacionado"`
	Emisor               Emisor            `json:"emisor"`
	Receptor             Receptor          `json:"receptor"`
	CuerpoDocumento      []CuerpoDocumento `json:"cuerpoDocumento"`
	Resumen              Resumen           `json:"resumen"`
	Extension            Extension         `json:"extension"`
	Apendice             interface{}       `json:"apendice"`
}

// Identificador representa la estructura de un identificador
type Identificador struct {
	Version          string `json:"version"`
	Ambiente         string `json:"ambiente"`
	TipoDTE          string `json:"tipoDTE"`
	NumeroControl    string `json:"numeroControl"`
	CodigoGeneracion string `json:"codigoGeneracion"`
	TipoModelo       string `json:"tipoModelo"`
	TipoOperacion    string `json:"tipoOperacion"`
	TipoContingencia string `json:"tipoContingencia"`
	MotivoContin     string `json:"motivoContin"`
	FechEmi          string `json:"fechEmi"`
	HorEmi           string `json:"horEmi"`
	TipoMoneda       string `json:"tipoMoneda"`
}

// Emisor representa la estructura de un emisor
type Emisor struct {
	Nit                 string      `json:"nit"`
	Ncr                 string      `json:"ncr"`
	Nombre              string      `json:"nombre"`
	CodActividad        string      `json:"codActividad"`
	DescActividad       string      `json:"descActividad"`
	NombreComercial     string      `json:"nombreComercial"`
	TipoEstablecimiento string      `json:"tipoEstablecimiento"`
	Direccion           Direccion   `json:"direccion"`
	Telefono            string      `json:"telefono"`
	Correo              string      `json:"correo"`
	CodEstable          interface{} `json:"codEstable"`
	CodPuntoVenta       interface{} `json:"codPuntoVenta"`
	CodEstableMH        string      `json:"codEstableMH"`
	CodPuntoVentaMH     string      `json:"codPuntoVentaMH"`
}

// Direccion representa la estructura de una dirección
type Direccion struct {
	Departamento string `json:"departamento"`
	Municipio    string `json:"municipio"`
	Complemento  string `json:"complemento"`
}

// Receptor representa la estructura de un receptor
type Receptor struct {
	Ncr             string    `json:"ncr"`
	Nombre          string    `json:"nombre"`
	CodActividad    string    `json:"codActividad"`
	DescActividad   string    `json:"descActividad"`
	Direccion       Direccion `json:"direccion"`
	Telefono        string    `json:"telefono"`
	Correo          string    `json:"correo"`
	NombreComercial string    `json:"nombreComercial"`
	Nit             string    `json:"nit"`
}

// CuerpoDocumento representa la estructura de un cuerpo de documento
type CuerpoDocumento struct {
	ID           string      `json:"id,omitempty"`
	DTEID        string      `json:"dteID"`
	NumItem      int         `json:"numItem"`
	TipoItem     int         `json:"tipoItem"`
	Cantidad     int         `json:"cantidad"`
	Codigo       string      `json:"codigo"`
	UniMedida    int         `json:"uniMedida"`
	Descripcion  string      `json:"descripcion"`
	PrecioUni    float64     `json:"precioUni"`
	MontoDescu   float64     `json:"montoDescu"`
	CodTributo   interface{} `json:"codTributo"`
	VentaNoSuj   int         `json:"ventaNoSuj"`
	VentaExenta  int         `json:"ventaExenta"`
	VentaGravada int         `json:"ventaGravada"`
	Tributos     []string    `json:"tributos"`
	Psv          int         `json:"psv"`
	NoGravado    int         `json:"noGravado"`
}

// Resumen representa la estructura de un resumen
type Resumen struct {
	ID                  string      `json:"id,omitempty"`
	DTEID               string      `json:"dteID"`
	TotalNoSuj          int         `json:"totalNoSuj"`
	TotalExenta         int         `json:"totalExenta"`
	TotalGravada        int         `json:"totalGravada"`
	SubTotalVentas      int         `json:"subTotalVentas"`
	DescuNoSuj          int         `json:"descuNoSuj"`
	DescuExenta         int         `json:"descuExenta"`
	DescuGravada        int         `json:"descuGravada"`
	PorcentajeDescuento int         `json:"porcentajeDescuento"`
	TotalDescu          int         `json:"totalDescu"`
	Tributos            []Tributo   `json:"tributos"`
	SubTotal            int         `json:"subTotal"`
	IvaPerci1           int         `json:"ivaPerci1"`
	IvaRete1            int         `json:"ivaRete1"`
	ReteRenta           int         `json:"reteRenta"`
	MontoTotalOperacion int         `json:"montoTotalOperacion"`
	TotalNoGravado      int         `json:"totalNoGravado"`
	TotalPagar          int         `json:"totalPagar"`
	TotalLetras         string      `json:"totalLetras"`
	SaldoFavor          int         `json:"saldoFavor"`
	CondicionOperacion  int         `json:"condicionOperacion"`
	Pagos               interface{} `json:"pagos"`
	NumPagoElectronico  interface{} `json:"numPagoElectronico"`
}

// Tributo representa la estructura de un tributo
type Tributo struct {
	ID          string `json:"id,omitempty"`
	ResumenID   string `json:"resumenID"`
	Codigo      string `json:"codigo"`
	Descripcion string `json:"descripcion"`
	Valor       int    `json:"valor"`
}

// Extension representa la estructura de una extensión
type Extension struct {
	ID            string `json:"id,omitempty"`
	DTEID         string `json:"dteID"`
	NombEntrega   string `json:"nombEntrega"`
	DocuEntrega   string `json:"docuEntrega"`
	NombRecibe    string `json:"nombRecibe"`
	DocuRecibe    string `json:"docuRecibe"`
	PlacaVehiculo string `json:"placaVehiculo"`
	Observaciones string `json:"observaciones"`
}

func init() {
	// Configurar el cliente de MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		panic("Error al crear el cliente de MongoDB: " + err.Error())
	}

	// Conectar al servidor de MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic("Error al conectar al servidor de MongoDB: " + err.Error())
	}

	// Seleccionar la base de datos y la colección
	database := client.Database("DTE_Reception")
	collection = database.Collection("dtes")
}
func obtenerTipoDTE() string {
	tipoDTE := []string{"01", "03", "04", "05", "06", "07", "08", "09", "11", "14", "15"}
	index := rand.Intn(len(tipoDTE))
	return tipoDTE[index]
}

func generarFechaAleatoria() string {
	min := time.Date(2023, time.June, 3, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	randomUnix := min + rand.Int63n(delta)

	randomTime := time.Unix(randomUnix, 0)
	return randomTime.Format("02/01/2006")
}

func insertarDatosPrueba() {
	for i := 0; i < 5; i++ {
		identificador := Identificador{
			Version:          "1.0",
			Ambiente:         "Produccion",
			TipoDTE:          obtenerTipoDTE(),
			NumeroControl:    fmt.Sprintf("Control%d", i),
			CodigoGeneracion: fmt.Sprintf("Generacion%d", i),
			FechEmi:          generarFechaAleatoria(), // Generar fecha aleatoria en el rango especificado
			HorEmi:           time.Now().Format("15:04:05"),
			TipoModelo:       "Electronic",
			TipoOperacion:    "01",
			TipoContingencia: "0",
			MotivoContin:     "",
			TipoMoneda:       "GTQ",
		}
		// Ejemplo de datos para Emisor
		emisor := Emisor{
			Nit:                 fmt.Sprintf("123456789-%d", i),
			Ncr:                 fmt.Sprintf("NCR%d", i),
			Nombre:              fmt.Sprintf("Empresa%d", i),
			CodActividad:        fmt.Sprintf("Act%d", i),
			DescActividad:       fmt.Sprintf("Actividad%d", i),
			NombreComercial:     fmt.Sprintf("Comercial%d", i),
			TipoEstablecimiento: fmt.Sprintf("Establecimiento%d", i),
			Direccion: Direccion{
				Departamento: fmt.Sprintf("Departamento%d", i),
				Municipio:    fmt.Sprintf("Municipio%d", i),
				Complemento:  fmt.Sprintf("Complemento%d", i),
			},
			Telefono: fmt.Sprintf("123456%d", i),
			Correo:   fmt.Sprintf("correo%d@example.com", i),
			// ... Resto de los campos según tu estructura Emisor
		}

		// Ejemplo de datos para Receptor
		receptor := Receptor{
			Ncr:           fmt.Sprintf("789012345-%d", i),
			Nombre:        fmt.Sprintf("Cliente%d", i),
			CodActividad:  fmt.Sprintf("ActCliente%d", i),
			DescActividad: fmt.Sprintf("ActividadCliente%d", i),
			Direccion: Direccion{
				Departamento: fmt.Sprintf("DepartamentoCliente%d", i),
				Municipio:    fmt.Sprintf("MunicipioCliente%d", i),
				Complemento:  fmt.Sprintf("ComplementoCliente%d", i),
			},
			Telefono: fmt.Sprintf("789012%d", i),
			Correo:   fmt.Sprintf("cliente%d@example.com", i),
			Nit:      fmt.Sprintf("NITCliente%d", i),
			// ... Resto de los campos según tu estructura Receptor
		}

		// Ejemplo de datos para CuerpoDocumento
		cuerpoDocumento := CuerpoDocumento{
			DTEID:       fmt.Sprintf("ID%d", i),
			NumItem:     i + 1,
			TipoItem:    i % 2, // Ejemplo de alternar entre 0 y 1
			Cantidad:    10 + i,
			Codigo:      fmt.Sprintf("Codigo%d", i),
			UniMedida:   1,
			Descripcion: fmt.Sprintf("Descripción%d", i),
			PrecioUni:   float64(50 + i),
			MontoDescu:  float64(i * 5),
			// ... Resto de los campos según tu estructura CuerpoDocumento
		}

		// Ejemplo de datos para Resumen
		resumen := Resumen{
			DTEID:               fmt.Sprintf("ID%d", i),
			TotalNoSuj:          1000 + i,
			TotalExenta:         500 + i,
			TotalGravada:        1500 + i,
			SubTotalVentas:      2000 + i,
			DescuNoSuj:          50 + i,
			DescuExenta:         20 + i,
			DescuGravada:        70 + i,
			PorcentajeDescuento: 5,
			TotalDescu:          140 + i,
			SubTotal:            1860 + i,
			IvaPerci1:           12,
			IvaRete1:            5,
			ReteRenta:           8,
			MontoTotalOperacion: 2000 + i,
			TotalNoGravado:      500 + i,
			TotalPagar:          1500 + i,
			TotalLetras:         "Mil Quinientos",
			SaldoFavor:          100,
			CondicionOperacion:  1,

			// ... Resto de los campos según tu estructura Resumen
		}

		// Ejemplo de datos para Tributo (puedes ajustar según tu estructura)
		tributo := Tributo{
			ResumenID:   resumen.ID,
			Codigo:      fmt.Sprintf("Tributo%d", i),
			Descripcion: fmt.Sprintf("DescripciónTributo%d", i),
			Valor:       10 + i,
		}

		// Asignar el tributo al resumen
		resumen.Tributos = append(resumen.Tributos, tributo)

		// Ejemplo de datos para Extension
		extension := Extension{
			DTEID:         fmt.Sprintf("ID%d", i),
			NombEntrega:   fmt.Sprintf("Entrega%d", i),
			DocuEntrega:   fmt.Sprintf("DocEntrega%d", i),
			NombRecibe:    fmt.Sprintf("Recibe%d", i),
			DocuRecibe:    fmt.Sprintf("DocRecibe%d", i),
			PlacaVehiculo: fmt.Sprintf("Placa%d", i),
			Observaciones: fmt.Sprintf("Observaciones%d", i),
			// ... Resto de los campos según tu estructura Extension
		}

		// Crear DTE con las estructuras anteriores
		dte := DTE{
			Identificador:   identificador,
			Emisor:          emisor,
			Receptor:        receptor,
			CuerpoDocumento: []CuerpoDocumento{cuerpoDocumento},
			Resumen:         resumen,
			Extension:       extension,
			Apendice:        nil, // Puedes ajustar según tu estructura
		}

		// Insertar en MongoDB
		_, err := collection.InsertOne(context.Background(), dte)
		if err != nil {
			fmt.Println("Error al insertar datos de prueba:", err)
		}
	}
}
func main() {
	r := gin.Default()

	// Añade un nuevo endpoint para obtener los documentos almacenadosb
	r.GET("/dte_search", func(c *gin.Context) {
		var dtes []DTE

		// Consultar MongoDB para obtener todos los documentos
		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error al mostrar los documentos": err.Error()})
			return
		}
		defer cursor.Close(context.Background())

		// Iterar sobre los documentos y agregarlos al slice
		for cursor.Next(context.Background()) {
			var dte DTE
			if err := cursor.Decode(&dte); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			dtes = append(dtes, dte)
		}

		c.JSON(http.StatusOK, dtes)
	})

	insertarDatosPrueba()
	r.Run(":8080")
}
