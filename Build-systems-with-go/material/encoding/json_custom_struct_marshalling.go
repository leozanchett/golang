package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type User struct {
	UserId   string `json:"user_id,omitempty"`
	Score    int    `json:"score,omitempty"`
	Password string `json:"password,omitempty"`
}

func main() {
	userA := User{"Gopher", 1000, "admin"}
	userB := User{"BigJ", 10, "1234"}
	userC := User{UserId: "GGBoom", Password: "1111"}

	db := []User{userA, userB, userC}
	dbJson, err := json.Marshal(&db)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dbJson))

	var recovered []User
	err = json.Unmarshal(dbJson, &recovered)
	if err != nil {
		panic(err)
	}
	fmt.Println(recovered)

	var indented bytes.Buffer
	err = json.Indent(&indented, dbJson, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(indented.String())

	// acessing the fields of a struct
	fmt.Println(recovered[0].UserId)
}
