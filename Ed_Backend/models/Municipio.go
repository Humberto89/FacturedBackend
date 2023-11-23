package models

import "github.com/jinzhu/gorm"

type Municipio struct {
	gorm.Model
	Nombre         string `gorm:"not null"`
	DepartamentoID uint   `gorm:"foreignKey:DepartamentoID"`
}
