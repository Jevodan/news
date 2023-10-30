package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Rss           []string `json:"rss"`
	RequestPeriod int      `json:"request_period"`
}

// Получаем ссылки на RSS новости и период обхода из конфигурационного json файла
func GetContentFromJson() ([]string, int, error) {
	var config Config
	file, err := os.Open("../../config/news.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return nil, 0, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return nil, 0, err
	}
	return config.Rss, config.RequestPeriod, nil
}
