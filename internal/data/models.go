package data

type Post struct {
	ID      uint64 `gorm:"primaryKey"`
	Title   string
	Author  string
	Date    string
	Content string
}
