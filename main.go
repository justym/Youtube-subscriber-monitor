package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Youtube subscriber monitor")
	handler := Handler{}
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/stats", handler.StatsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


	

