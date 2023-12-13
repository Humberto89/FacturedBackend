package models

import "github.com/jinzhu/gorm"

type Pais struct {
	gorm.Model
	Nombre        string         `gorm:"unique;not null"`
	Departamentos []Departamento `gorm:"ForeignKey:PaisID"`
}
