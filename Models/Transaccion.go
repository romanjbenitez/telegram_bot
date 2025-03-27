package Models

import (
	"time"
)

type Transaccion struct {
	ID               uint            `gorm:"primaryKey"`
	Cantidad         float64         `gorm:"not null"`
	Tipo             TipoTransaccion `gorm:"not null"`
	Fecha            time.Time       `gorm:"not null"`
	FormaDePago      FormaDePago     `gorm:"not null"`
	Categoria        string          `gorm:"not null"`
	Descripcion      string
	PlataformaDePago PlataformaDePago `gorm:"type:varchar(20);not null"`
	CreadoEn         time.Time        `gorm:"autoCreateTime"`
}
