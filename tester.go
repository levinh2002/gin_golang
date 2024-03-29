package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/videos"
	method := "GET"

	payload := strings.NewReader(`{` + " " + ` "title": "t fdsfsfdsfdd1",` + " " + ` "description": "d fdsfdsf1",` + " " + ` "url": "url fdsfs1"` + " " + ` }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
