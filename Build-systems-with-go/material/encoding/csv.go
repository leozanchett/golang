package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	in := `user_id,score,password
"Gopher",1000,"admin"
"BigJ",10,"1234"
"GGBoom",,"1111"`

	out := [][]string{
		{"user_id", "score", "password"},
		{"Gopher", "1000", "admin"},
		{"BigJ", "10", "1234"},
		{"GGBoom", "", "1111"},
	}

	r := csv.NewReader(strings.NewReader(in))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)
	for _, rec := range out {
		if err := writer.Write(rec); err != nil {
			panic(err)
		}
	}
	writer.Flush()
}
