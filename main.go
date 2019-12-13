package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justym/youtube-subscriber-monitor/websocket"
	"github.com/justym/youtube-subscriber-monitor/youtube"
)

func main() {
	//SetRouters()

	item, err := youtube.GetSubscribers()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v\n", item)
	fmt.Println("Youtube subscriber monitor")
	SetRouters()
}

//SetRouters : Router setting
func SetRouters() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/stats", StatsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//HomeHandler : Display sentence at '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Youtube subscriber monitor")
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)
}
