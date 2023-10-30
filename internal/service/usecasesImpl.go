package service

import (
	"log"
	"news/internal/data"
	"news/models"
)

type UsecasesImpl struct {
	db data.DataBase
}

func New(db data.DataBase) *UsecasesImpl {
	return &UsecasesImpl{
		db: db,
	}
}

// CreateNew имплементация Usecases.
func (u *UsecasesImpl) CreateNew(news []models.Post) error {
	for _, new := range news {
		err := u.db.AddNews(new)
		if err != nil {
			log.Printf("Ошибка: %v", err)
			return err
		}
	}
	return nil
}

// GetNews имплементация Usecases.
func (u *UsecasesImpl) GetNews(id int) ([]models.Post, error) {
	res, err := u.db.News(id)
	if err != nil {
		log.Printf("ОШибка получения новостей %v", err)
	}
	return res, nil
}
