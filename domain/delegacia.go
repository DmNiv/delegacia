package domain

import "github.com/google/uuid"

type Delegacia struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key"`
	Nome      string    `json:"nome"`
	Endereco  string    `json:"endereco"`
	DiaTodo   bool      `json:"diaTodo"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	MapUrl    string    `json:"mapUrl"`
	Telefone  string    `json:"telefone"`
}
