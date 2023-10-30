package main

import (
	"log"
	"net/http"
	"news/config"
	api "news/internal/api/http"
	"news/internal/api/http/handlers"
	"news/internal/data"
	"news/internal/service"
	"time"
)

func main() {
	repo := data.NewSQL()
	usecases := service.New(repo)
	r := api.SetupRouter(usecases)
	addNews(usecases)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Ошибка при запуске HTTP-сервера: ", err)
	}
}

// Функция запускает горутины обхода RSS лент
func addNews(u service.Usecases) {
	rssLenta, period, err := config.GetContentFromJson()
	if err != nil {
		log.Fatal("Ошибка чтения RSS ленты", err)
	}

	for _, val := range rssLenta {
		go func(val string) {
			for {
				newsHandler := handlers.NewHandler(u)
				err := newsHandler.NewsCreate(val)
				if err != nil {
					log.Fatal("Ошибка создания новости", err)
				}
				time.Sleep(time.Minute * time.Duration(period))
			}
		}(val)
	}
}
