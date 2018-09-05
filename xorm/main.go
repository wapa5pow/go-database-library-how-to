package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

type User struct {
	ID   int64  `json:"id" xorm:"'id'"`
	Name string `json:"name" xorm:"'nickname'"`
}

func main() {
	// FIXME: it does not work
	var err error
	engine, err = xorm.NewEngine("postgres", "user=postgres dbname=postgres sslmode=disable password=password")
	if err != nil {
		panic(err)
	}

	err = engine.Sync2(new(User))
	if err != nil {
		panic(err)
	}

	user := &User{Name: "Ishida"}
	affected, err := engine.Insert(&user)
	if err != nil {
		panic(err)
	}
	j, _ := json.Marshal(affected)
	fmt.Println(string(j))
}
