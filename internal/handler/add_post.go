package handler

import (
	"candy/internal/data"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (h Handler) AddPostHandler(w http.ResponseWriter, r *http.Request) error {
	var post *data.Post

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, post); err != nil {
		return err
	}

	return h.DB.Add(post)
}
