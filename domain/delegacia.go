package domain

import "github.com/google/uuid"

type Delegacia struct {
	Id         uuid.UUID `json:"id" gorm:"primary_key"`
	Nome       string    `json:"nome"`
	Endereco   string    `json:"endereco"`
	Horario24h bool      `json:"horario24h"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
}
