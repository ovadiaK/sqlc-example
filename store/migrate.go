package store

import (
	"fmt"
	"log"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() {
	m, err := migrate.New(
		"file://../store/migrations",
		"postgresql://tutorial:abc@172.17.0.2:5432/tutorial?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("migrations started?")
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
