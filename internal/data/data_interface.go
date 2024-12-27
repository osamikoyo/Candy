package data

import "sync"

type PostStorage interface {
	GetByTitle(title string) ([]Post, error)
	Add(post *Post) error
	Get() ([]Post, error)
	Last(ch chan Post, cher chan error, wg *sync.WaitGroup, mux *sync.Mutex)
}
