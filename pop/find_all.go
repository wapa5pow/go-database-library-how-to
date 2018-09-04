package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin/json"
	"github.com/gobuffalo/pop"
	"github.com/wapa5pow/go-database-library-how-to/pop/models"
)

func main() {
	conn, err := pop.Connect("development")
	if err != nil {
		log.Fatal(err)
	}

	users := []*models.User{}
	err = conn.All(&users)
	if err != nil {
		log.Fatal(err)
	}

	j, _ := json.Marshal(users)
	fmt.Println(string(j))
}
