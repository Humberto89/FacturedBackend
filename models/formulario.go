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

type ExtranjeroType string

const (
	Si ExtranjeroType = "Si"
	No ExtranjeroType = "No"
)

type ExtentoType string

const (
	Sies ExtentoType = "Si"
	Noes ExtentoType = "No"
)

type Formulario struct {
	gorm.Model
	ClienteType          ClienteType `gorm:"not null"`
	Nombres              string      `gorm:"not null"`
	Apellidos            string      `gorm:"not null"`
	Extranjero           ExtranjeroType
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
	RecintoFiscal        string
	ActividadEconomica   string
	TituloRemisionBienes string
	ContribuyenteType    ContribuyenteType `gorm:"not null"`
	ExencionIVA          ExtentoType
}
