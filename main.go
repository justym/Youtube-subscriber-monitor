package main

import (
	"fmt"
	"github.com/justym/youtube-subscriber-monitor/youtube"
	"log"
	"net/http"
)

func main() {
	SetRouters()

	item, err := youtube.GetSubscribers()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("%+v\n", item)
	fmt.Println("Youtube subscriber monitor")
}

//SetRouters : Router setting
func SetRouters() {
	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//HomeHandler : Display sentence at '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Youtube subscriber monitor")
}
