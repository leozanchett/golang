package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	bodyRequest := []byte(`"user":"john", "email":"john@gmail.com"`)
	bufferBody := bytes.NewBuffer(bodyRequest)
	const url_post_data = "https://httpbin.org/post"
	resp, err := http.Post(url_post_data, "application/json", bufferBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	bodyAnswer := bufio.NewScanner(resp.Body)
	for bodyAnswer.Scan() {
		fmt.Println(bodyAnswer.Text())
	}

	// POST FORM
	resp2, err2 := http.PostForm(url_post_data, url.Values{"user": {"john"}, "email": {"john@gmail.com"}})
	if err2 != nil {
		panic(err2)
	}
	defer resp2.Body.Close()
	fmt.Println(resp2.Status)
	bodyAnswer2 := bufio.NewScanner(resp2.Body)
	for bodyAnswer2.Scan() {
		fmt.Println(bodyAnswer2.Text())
	}

}
