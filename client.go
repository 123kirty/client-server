package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func checkErr() {
	var (
		err error
	)

	if err != nil {
		log.Println(err)
	}
}

func RunClient() {
	resp, _ := http.Get("http://localhost:8080/")
	checkErr()
	val, _ := ioutil.ReadAll(resp.Body)
	checkErr()
	fmt.Println(string(val))
}

func main() {
	RunClient()
}
