package controllers

import(
	"net/http"
	"app/utils"
	"app/models"
	"github.com/gorilla/mux"
)

// Получить все транзакции, с возможностью задать фильтры
func GetTransactions(w http.ResponseWriter, r *http.Request){
	user_id := r.Context().Value("user_id").(int)
	query := r.URL.Query()
    transaction := models.Transaction{UserId:user_id}
    t := transaction.GetAll(query)
    utils.Response(w, t)
}

// Получить урл на все транзакции в csv файле, с возможностью задать фильтры
func GetTransactionsCsv(w http.ResponseWriter, r *http.Request){
	user_id := r.Context().Value("user_id").(int)
	query := r.URL.Query()
    transaction := models.Transaction{UserId:user_id}
    t := transaction.GetAllCsv(query)
    utils.Response(w, t)
}

// Получить одну транзакцию по id
func GetTransaction(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	user_id := r.Context().Value("user_id").(int)
	transaction := models.Transaction{UserId:user_id}
	t := transaction.GetOne(vars["id"])
	utils.Response(w, t)
}

