package router

import (
	"app/controllers"
	"app/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// Функция инициализирующая роутер

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.JwtAuthentication)                                                                  // Подключаем прокладку для проверки токена
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/")))) // Указываем статическую папку
	r.HandleFunc("/auth/login", controllers.LoginHandle).Methods("POST")
	r.HandleFunc("/auth/registration", controllers.RegistrationHandle).Methods("POST")
	r.HandleFunc("/sms/send", controllers.SendSms).Methods("POST")
	r.HandleFunc("/sms/check", controllers.CheckSms).Methods("POST")
	r.HandleFunc("/transactions/get", controllers.GetTransactions).Methods("GET")
	r.HandleFunc("/transactions/csv", controllers.GetTransactionsCsv).Methods("GET")
	r.HandleFunc("/transactions/get/{id}", controllers.GetTransaction).Methods("GET")
	r.HandleFunc("/user/update", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/upload/picture", controllers.ProfilePicture).Methods("POST")
	return r
}
