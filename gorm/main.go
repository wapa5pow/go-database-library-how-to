package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name string
}

func main() {

	db, err := gorm.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=password")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)

	var newUser = &User{Name: "Ishida"}
	newDB := db.AutoMigrate(&User{})
	if newDB.Error != nil {
		panic(newDB.Error)
	}
	newDB = db.Create(newUser)
	if newDB.Error != nil {
		panic(newDB.Error)
	}

	j, _ := json.Marshal(newUser)
	fmt.Println(string(j))

	var user User
	newDB = db.First(&user, 1)
	if newDB.Error != nil {
		panic(newDB.Error)
	}

	j, _ = json.Marshal(user)
	fmt.Println(string(j))

	newDB = db.Delete(&user)
	if newDB.Error != nil {
		panic(newDB.Error)
	}
}
