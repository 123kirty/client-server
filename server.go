package main

import (
    "fmt"
    "net/http"
    "log"
)

func Handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"bro present!")
}

func RunServer(){
  http.HandleFunc("/",Handler)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func main(){
    RunServer()
}

