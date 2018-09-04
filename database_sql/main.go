package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	type User struct {
		Id   int
		Name string
	}

	db, _ := sql.Open("postgres",
		"user=postgres password=password dbname=postgres sslmode=disable")

	rows, err := db.Query(`SELECT * FROM "users"`)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	users := []User{}
	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}
	fmt.Println(users)
}
