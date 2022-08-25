package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ovadiaK/sqlc-example/store"
	"log"
	"reflect"

	"github.com/ovadiaK/sqlc-example/store/tutorial"

	_ "github.com/lib/pq"
)

func run() error {
	ctx := context.Background()

	db, err := sql.Open("postgres", "postgresql://tutorial:abc@172.17.0.2:5432/tutorial?sslmode=disable")
	if err != nil {
		return err
	}
	fmt.Println("store connected")

	store.Migrate()

	queries := tutorial.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	fmt.Println("hello world")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
