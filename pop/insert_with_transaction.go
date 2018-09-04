package main

import (
	"log"

	"github.com/gobuffalo/pop"
	"github.com/wapa5pow/go-database-library-how-to/pop/models"
)

func main() {
	conn, err := pop.Connect("development")
	if err != nil {
		log.Fatal(err)
	}

	conn, err = conn.NewTransaction()
	if err != nil {
		log.Fatal(err)
	}
	user := models.User{Name: "Ishida"}
	conn.Create(&user)
	conn.TX.Commit()
}