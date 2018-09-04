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
	err = conn.Create(&user)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.TX.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
