package handler

import (
	"candy/internal/data"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

func broadcastNewPosts(post *data.Post) error {
	body, err := json.Marshal(post)
	if err != nil {
		return err
	}
	for client := range clients {
		err = client.WriteJSON(&body)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h Handler) ListenNewPosts(errors chan error) {
	var mux *sync.Mutex
	var wg *sync.WaitGroup
	var ch chan data.Post
	var cherr chan error

	go h.DB.Last(ch, cherr, wg, mux)

	for {
		wg.Add(1)

		select {
		case post := <-ch:
			err := broadcastNewPosts(&post)
			errors <- err
		case err := <-cherr:
			errors <- err
		}
	}
}

func (h Handler) GetAllHandler(w http.ResponseWriter, r *http.Request) error {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	defer func(ws *websocket.Conn) {
		_ = ws.Close()
	}(ws)
	clients[ws] = true

	var errs chan error

	h.ListenNewPosts(errs)

	err = <-errs
	if err != nil {
		return err
	}
	select {}

	return nil
}
