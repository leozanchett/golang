package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	UserId   string `pretty:"upper"`
	Email    string `pretty:"lower"`
	Password string `pretty:"lower"`
}

type Record struct {
	Name    string `pretty:"lower" json:"name"`
	Surname string `pretty:"upper" json:"surname"`
	Age     int    `pretty:"other" json:"age"`
}

func Marshal(input interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	buffer.WriteString("{")
	for i := 0; i < t.NumField(); i++ {
		encodedField, err := encodeField(t.Field(i), v.Field(i))
		if err != nil {
			return nil, err
		}
		if len(encodedField) != 0 {
			if i > 0 && i <= t.NumField() {
				buffer.WriteString(", ")
			}
			buffer.WriteString(encodedField)
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func encodeField(field reflect.StructField, value reflect.Value) (string, error) {
	if field.PkgPath != "" {
		return "", nil
	}
	if field.Type.Kind() != reflect.String {
		return "", nil
	}
	tag, found := field.Tag.Lookup("pretty")
	if !found {
		return "", nil
	}
	result := field.Name + ":"
	var err error = nil
	switch tag {
	case "upper":
		result += strings.ToUpper(value.String())
	case "lower":
		result += strings.ToLower(value.String())
	default:
		err = errors.New("unknown tag")
	}
	if err != nil {
		return "", err
	}

	return result, nil
}

func main() {
	u := User{"John", "John@Gmail.com", "admin"}
	marSer, _ := Marshal(u)
	fmt.Println("pretty user: ", string(marSer))

}
