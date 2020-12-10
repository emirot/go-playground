package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func printMySite() {
	resp, _ := http.Get("https://www.nolanemirot.com")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}

func main() {
	printMySite()
}
