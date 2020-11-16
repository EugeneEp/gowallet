package main

import (
	"net/http"
	"app/router"
	"github.com/rs/cors"
	"app/models"
)

func main(){
	// Включаем кросс доменность
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	models.Init() // Инициализируем соединение с бд
	defer models.DB.Close() // Запускаем закрытие соединения с бд, как отложенное событие
	r := router.NewRouter() // Подключаем роутер
	http.Handle("/", r)
	http.ListenAndServe(":8080", c.Handler(r)) // Запускаем сервер 
}