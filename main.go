package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//SetRouters()
	fmt.Println("Youtube subscriber monitor")
	SetRouters()
}

//SetRouters : Router setting
func SetRouters() {
	handler := Handler{}
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/stats", handler.StatsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
