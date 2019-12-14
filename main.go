package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := Handler{}
	http.HandleFunc("/stats", handler.StatsHandler)
	staticHandler := http.FileServer(http.Dir("./"))
	http.Handle("/", staticHandler)

	fmt.Println("Youtube subscriber monitor")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
