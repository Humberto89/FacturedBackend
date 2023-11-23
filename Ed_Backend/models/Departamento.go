package models

import "github.com/jinzhu/gorm"

type Departamento struct {
	gorm.Model
	Nombre     string      `gorm:"not null"`
	PaisID     uint        `gorm:"index"`
	Municipios []Municipio `gorm:"ForeignKey:DepartamentoID"`
}
