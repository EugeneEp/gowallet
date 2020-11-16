package utils

import (
	"net/http"
	"encoding/json"
)

// Функция отправки ответа

func Response(w http.ResponseWriter, m map[string]interface{}) {
	js, err := json.Marshal(m)
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if m["success"] == false {
		w.WriteHeader(http.StatusBadRequest)
	}else{
		w.WriteHeader(http.StatusOK)
	}
	
	w.Write(js)
}