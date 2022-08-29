package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.Header["Content-Type"])

	buf := bufio.NewScanner(resp.Body)
	for buf.Scan() {
		fmt.Println(buf.Text())
	}

}
