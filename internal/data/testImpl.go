package data

import (
	"news/models"
	"sync"
)

// Тестовая БД в мапе
type DB struct {
	s     sync.Mutex
	id    int
	store map[int]models.Post
}

func New() *DB {
	db := DB{
		id:    0,
		store: map[int]models.Post{},
	}
	return &db
}

func (db *DB) AddNews(new models.Post) error {
	db.s.Lock()
	defer db.s.Unlock()
	db.store[db.id] = new
	db.id++
	return nil
}

func (db *DB) News(num int) ([]models.Post, error) {
	db.s.Lock()
	defer db.s.Unlock()
	var news []models.Post
	for i := db.id - 1; i >= db.id-num; i-- {
		news = append(news, db.store[i])
	}
	return news, nil
}
