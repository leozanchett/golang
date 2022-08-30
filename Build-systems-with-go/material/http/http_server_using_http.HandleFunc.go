package main

import (
	"fmt"
	"net/http"
)

func info(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		fmt.Println(name, headers)
	}
	w.Write([]byte("Perfect !!!"))
	return
}

func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe(":8090", nil)
}
