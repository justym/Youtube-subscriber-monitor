package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justym/youtube-subscriber-monitor/websocket"
)

//Handler is struct of http Handler
type Handler struct{}

//StatsHandler handle stats
func (h *Handler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)
}
