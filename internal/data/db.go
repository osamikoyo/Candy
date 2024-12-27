package data

import (
	"candy/pkg/loger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func New() *Database {
	db, err := gorm.Open(sqlite.Open("storage/main.db"))
	if err != nil {
		loger.New().Error().Err(err)
	}

	return &Database{db}
}

func (d *Database) GetByTitle(title string) ([]Post, error) {
	var posts []Post

	if err := d.Where(Post{
		Title: title,
	}).Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (d *Database) Add(post *Post) error {
	return post.Save(d)
}

func (d *Database) Get() ([]Post, error) {
	var posts []Post

	err := d.Find(&posts).Error
	return posts, err
}
