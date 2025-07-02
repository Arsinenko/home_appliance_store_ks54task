package pkg

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func DatabaseInit() *pgx.Conn {
	//init postgres database using pgx
	db, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/home_appliance_store")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Println("Connected to database")
	return db
}
