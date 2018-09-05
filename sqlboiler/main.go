package main

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/wapa5pow/go-database-library-how-to/sqlboiler/models"
)

func main() {
	// Open handle to database like normal
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=password")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Commit()

	var p1 models.Pilot
	p1.Name = "Larry"
	err = p1.Insert(ctx, db, boil.Whitelist("id", "name"))
	if err != nil {
		panic(err)
	}
}
