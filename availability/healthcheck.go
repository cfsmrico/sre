package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("www.fitcalcs.life")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() // Ensure the response body is closed

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println(string(body))
}
