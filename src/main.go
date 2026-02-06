package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("listening at :80")
	log.Fatal(http.ListenAndServe(":80", router()))
}
