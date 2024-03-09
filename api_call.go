package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiUrl := "https://bookbeta.com/api/v1"

	response, err := http.Get(apiUrl)

	if err != nil {
		fmt.Println("Error making request: ", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: ", err)
		return
	}
	fmt.Println(string(body))
}
