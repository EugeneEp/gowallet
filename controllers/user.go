package controllers

import(
	"net/http"
	"app/utils"
	"app/models"
	"encoding/json"
)

// Получить токен
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Response(w, utils.ErrServerError)
	}
	user.Init()
	login := user.Login()
	utils.Response(w, login)
}

// Создать нового пользователя
func RegistrationHandle(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Response(w, utils.ErrServerError)
	}
	user.Init()
	reg := user.CreateUser()
	utils.Response(w, reg)
}

// Обновить данные пользователя
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Response(w, utils.ErrServerError)
	}
	update := user.UpdateUser(user_id)
	utils.Response(w, update)
}

// Загрузить фото профиля
func ProfilePicture(w http.ResponseWriter, r *http.Request){
	user_id := r.Context().Value("user_id").(int)
	user := models.User{}
	file := user.ProfilePicture(r, user_id)
	utils.Response(w, file)
}