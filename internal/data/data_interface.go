package data

type PostStorage interface {
	GetByTitle(title string) ([]Post, error)
	Add(post *Post) error
	Get() ([]Post, error)
}
