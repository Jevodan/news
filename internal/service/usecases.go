package service

import "news/models"

type Usecases interface {
	GetNews(id int) ([]models.Post, error)
	CreateNew(news []models.Post) error
}
