package main

import (
	"fmt"

	"github.com/gin-gonic/gin/json"

	"github.com/gobuffalo/pop"
	"github.com/wapa5pow/go-database-library-how-to/pop/models"
)

func main() {
	conn, err := pop.Connect("development")
	if err != nil {
		panic(err)
	}

	user := models.User{Name: "Ishida"}
	err = conn.Create(&user)
	if err != nil {
		panic(err)
	}

	books := []models.Book{
		{Title: "Book1", UserID: user.ID},
		{Title: "Book2", UserID: user.ID},
	}
	err = conn.Create(&books)
	if err != nil {
		panic(err)
	}

	matchUser := models.User{}
	err = conn.Eager().Find(&matchUser, user.ID)
	if err != nil {
		panic(err)
	}
	j, _ := json.Marshal(matchUser.Books)
	fmt.Println(string(j))
}
