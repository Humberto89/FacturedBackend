package models

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
}
