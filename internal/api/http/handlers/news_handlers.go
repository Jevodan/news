package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"news/internal/service"
	"news/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
)

type API struct {
	service service.Usecases
}

func NewHandler(service service.Usecases) *API {
	return &API{
		service: service,
	}
}

// метод для вывода новостей, для обработчика /news/{count}
func (api *API) NewsHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["count"]
	id, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // ошибка клиентская, передача айди серверу некорректна
		return
	}
	n, err := api.service.GetNews(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(n)
	w.WriteHeader(http.StatusOK)
}

// метод для создания новостей
func (api *API) NewsCreate(rssUrl string) error {
	var news []models.Post
	dateFormat := "Mon, 02 Jan 2006 15:04:05 -0700"
	parser := gofeed.NewParser()
	// Парсим RSS-ленту.
	feed, err := parser.ParseURL(rssUrl)
	if err != nil {
		log.Fatal("Ошибка при чтении RSS-ленты:", err)
		return err
	}

	for _, item := range feed.Items {
		t, err := time.Parse(dateFormat, item.Published)
		if err != nil {
			fmt.Println("Ошибка парсинга даты", err)
			return err
		}
		tUnix := t.Unix()
		new := models.Post{
			Title: item.Title, Description: item.Description, PubDate: tUnix, Link: item.Link,
		}
		news = append(news, new)
	}

	err = api.service.CreateNew(news)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
