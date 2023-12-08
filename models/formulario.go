package models

import "github.com/jinzhu/gorm"

// PersonaType representa el tipo de persona (natural o jurídica)
type ClienteType string

const (
	TipoNatural  ClienteType = "Natural"
	TipoJuridico ClienteType = "Juridico"
)

type ContribuyenteType string

const (
	OtrosContribuyentes  ContribuyenteType = "Otros Contribuyentes"
	MedianoContribuyente ContribuyenteType = "Mediano Contribuyente"
	GrandeContribuyente  ContribuyenteType = "Grande Contribuyente"
)

type titulo string

const (
	Deposito     titulo = "Deposito"
	Propiedad    titulo = "Propiedad"
	Consignacion titulo = "Consignación"
	Traslado     titulo = "Traslado"
	Otros        titulo = "Otros"
)

type recinto string

const (
	MarítimadeAcajutla   recinto = "Marítima de Acajutla"
	AéreaMonseñor        recinto = "Aérea Monseñor Óscar Arnulfo Romero"
	TerrestreLasChinamas recinto = "Terrestre Las Chinamas"
	TerrestreLaHachadura recinto = "Terrestre La Hachadura"
	TerrestreSantaAna    recinto = "Terrestre Santa Ana"
)

type Formulario struct {
	gorm.Model
	Dui                  string      `gorm:"unique"`
	NIT                  string      `gorm:"unique;not null"`
	ClienteType          ClienteType `gorm:"not null"`
	Nombres              string      `gorm:"not null"`
	Apellidos            string      `gorm:"not null"`
	ExtranjeroType       string
	NRC                  string
	CodigoCliente        string
	PaisID               uint         `gorm:"column:pais_id"`
	Pais                 Pais         `gorm:"foreignKey:id;references:PaisID"`
	DepartamentoID       uint         `gorm:"column:departamento_id"`
	Departamento         Departamento `gorm:"foreignKey:id;references:DepartamentoID"`
	MunicipioID          uint         `gorm:"column:municipio_id"`
	Municipio            Municipio    `gorm:"foreignKey:id;references:MunicipioID"`
	Direccion            string       `gorm:"not null"`
	Telefono             string
	Correo               string
	RecintoFiscal        recinto
	ActividadEconomica   string
	TituloRemisionBienes titulo
	ContribuyenteType    ContribuyenteType `gorm:"not null"`
	ExtentoType          string
	RazonSocial          string `gorm:"not null"`
	NombreComercial      string `gorm:"not null"`
}
