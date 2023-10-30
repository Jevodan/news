package data

import (
	"database/sql"
	"fmt"
	"log"
	"news/models"
	"sync"

	"github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
	s  sync.Mutex
}

// создание объекта БД Postgres
func NewSQL() *Storage {
	// Строка подключения к базе данных PostgreSQL.
	connStr := "user=leka password=serpent dbname=news host=localhost sslmode=disable"
	// Открываем соединение с базой данных.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка", err)
	}
	fmt.Println("Успешное подключение к PostgreSQL!")
	storage := Storage{
		db: db,
	}
	return &storage
}

// Добавление новости в БД
func (s *Storage) AddNews(new models.Post) error {
	s.s.Lock()
	defer s.s.Unlock()

	sql := "INSERT INTO news (title, description, publish, link) VALUES ($1,$2,$3,$4)"
	// Выполняем SQL-запрос для вставки данных
	_, err := s.db.Exec(sql, new.Title, new.Description, new.PubDate, new.Link)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code.Name() == "unique_violation" {
				fmt.Printf("Дубликат  %v \n", err)
			}
		} else {
			fmt.Printf("Ошибка записи в БД %v %v", err, sql)
			return err
		}
	}
	return nil
}

// Получение списка новостей
func (s *Storage) News(num int) ([]models.Post, error) {
	s.s.Lock()
	defer s.s.Unlock()
	var news []models.Post
	rows, err := s.db.Query("SELECT title,description,publish,link FROM news ORDER BY id DESC LIMIT $1", num)
	if err != nil {
		fmt.Printf("Ошибка получения данных")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var new models.Post
		if err := rows.Scan(&new.Title, &new.Description, &new.PubDate, &new.Link); err != nil {
			log.Fatalf("Ошибка записи данных в структуру %v", err)
		}
		news = append(news, new)
	}

	return news, nil
}
