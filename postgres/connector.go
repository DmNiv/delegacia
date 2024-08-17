package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	fmt.Println("Connecting to PostgreSQL...")

	dsn := "host=localhost user=oficina-futuro password=delegaciafacil dbname=delegacia-facil port=5432 sslmode=disable search_path=delegacia_facil"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected successfully")
	return DB
}
