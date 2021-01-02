package main

import (
	"app/models"
	"app/router"
	"net/http"
	"github.com/rs/cors"
	"github.com/joho/godotenv"
	"log"
)

// Загружаем конфигурационный файл
func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {

	models.Init()                              // Инициализируем соединение с бд
	defer models.DB.Close()
	
	models.InitRedis()						   // Инициализируем соединение с Redis
	defer models.RDB.Close()

	// Включаем кросс доменность
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	r := router.NewRouter()                    // Подключаем роутер

	http.ListenAndServe(":8080", c.Handler(r)) // Запускаем сервер

}
