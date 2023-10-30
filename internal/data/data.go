package data

import "news/models"

type DataBase interface {
	AddNews(models.Post) error
	News(int) ([]models.Post, error)
}
