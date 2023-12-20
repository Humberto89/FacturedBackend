package models

import "time"

// Estructura para decodificar el documento JSON
type DTE struct {
	ID         string    `json:"_id"`
	Filename   string    `json:"filename"`
	Size       int       `json:"size"`
	UploadDate time.Time `json:"uploadDate"`
	Data       Datos     `json:"data"`
}

type Datos struct {
	Identificacion Identificacion `json:"identificacion"`
	Resumen        Resumen        `json:"resumen"`
}

type Identificacion struct {
	Version          int         `json:"version"`
	Ambiente         string      `json:"ambiente"`
	TipoDte          string      `json:"tipoDte"`
	NumeroControl    string      `json:"numeroControl"`
	CodigoGeneracion string      `json:"codigoGeneracion"`
	TipoModelo       int         `json:"tipoModelo"`
	TipoOperacion    int         `json:"tipoOperacion"`
	TipoContingencia interface{} `json:"tipoContingencia"`
	MotivoContin     interface{} `json:"motivoContin"`
	FecEmi           string      `json:"fecEmi"`
	HorEmi           string      `json:"horEmi"`
	TipoMoneda       string      `json:"tipoMoneda"`
}

type Resumen struct {
	CondicionOperacion int `json:"condicionOperacion"`
}
