package model

type Local struct {
	ID       uint   `gorm:"primaryKey"`
	Nombre   string
	IP   string
	Numero   string
}

