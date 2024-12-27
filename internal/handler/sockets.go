package handler

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

func (h Handler) GetAllHandler(w http.ResponseWriter, r *http.Request) error {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	defer ws.Close()
	clients[ws] = true

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			delete(clients, ws)
			return err
		}

	}
}
