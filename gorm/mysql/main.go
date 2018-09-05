package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
}

func main() {

	db, err := gorm.Open("mysql", "username:password@/sample?charset=utf8&parseTime=True&loc=Local")
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
