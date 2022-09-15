package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserId   string `tagA:"valueA1" tagB:"valueA2"`
	Email    string `tagB:"value"`
	Password string `tagC:"v1 v2"`
}

func main() {
	u := User{"1", "2", "3"}
	T := reflect.TypeOf(u)

	fieldUserId := T.Field(0)
	t := fieldUserId.Tag
	fmt.Println("Struct field tag:", t)
	v, _ := t.Lookup("tagA")
	fmt.Printf("tagA: %s\n", v)

	fieldEmail, _ := T.FieldByName("Email")
	vEmail := fieldEmail.Tag.Get("tagB")
	fmt.Println("email tagB:", vEmail)

	fieldPassword, _ := T.FieldByName("Password")
	fmt.Printf("Password tags: [%s]\n", fieldPassword.Tag)
	fmt.Printf("Password tagC: [%s]\n", fieldPassword.Tag.Get("tagC"))
	fmt.Println(fieldPassword.Tag.Get("tagC"))

}
