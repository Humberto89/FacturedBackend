package models

import "time"

// Estructura para decodificar el documento JSON
type Documento struct {
	ID         string    `json:"_id"`
	Filename   string    `json:"filename"`
	Size       int       `json:"size"`
	Status     int       `json:"status"`
	UploadDate time.Time `json:"uploadDate"`
	Data       Ident     `json:"data"`
}

type Ident struct {
	Identificacion struct {
		Version          int    `json:"version"`
		Ambiente         string `json:"ambiente"`
		TipoDte          string `json:"tipoDte"`
		NumeroControl    string `json:"numeroControl"`
		CodigoGeneracion string `json:"codigoGeneracion"`
		TipoModelo       int    `json:"tipoModelo"`
		TipoOperacion    int    `json:"tipoOperacion"`
		TipoContingencia any    `json:"tipoContingencia"`
		MotivoContin     any    `json:"motivoContin"`
		FecEmi           string `json:"fecEmi"`
		HorEmi           string `json:"horEmi"`
		TipoMoneda       string `json:"tipoMoneda"`
	} `json:"identificacion"`
	Resumen struct {
		CondicionOperacion int `json:"condicionOperacion"`
	} `json:"resumen"`
}
