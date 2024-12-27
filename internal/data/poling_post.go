package data

import (
	"sync"
	"time"
)

func (d Database) Last(ch chan Post, cher chan error, wg *sync.WaitGroup, mux *sync.Mutex) {
	var latePostTitle string

	for {
		mux.Lock()
		var latepost Post

		if err := d.Order(&latepost).Error; err != nil {
			cher <- err
			time.Sleep(5 * time.Second)
			continue
		}

		if latepost.Title != latePostTitle {
			latePostTitle = latepost.Title
			ch <- latepost
		}

		mux.Unlock()

		time.Sleep(10 * time.Second)

		wg.Done()
	}
}
