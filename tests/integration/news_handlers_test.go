package handlers

import (
	"net/http"
	"net/http/httptest"
	api "news/internal/api/http"
	"news/internal/data"
	"news/internal/service"
	"news/models"
	"testing"
)

func TestAPI_NewsHandler(t *testing.T) {
	repo := data.New()
	new := models.Post{
		Title: "эта новость",
	}
	repo.AddNews(new)
	usecases := service.New(repo)

	server := httptest.NewServer(api.SetupRouter(usecases))
	defer server.Close()

	//  HTTP-клиент для отправки запросов на сервер.
	client := server.Client()

	//  HTTP-запрос,  отправляем на сервер.
	req, err := http.NewRequest("GET", server.URL+"/news/10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Отправьте запрос на сервер с помощью HTTP-клиента.
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус OK, получено: %v", resp.Status)
	}

}
