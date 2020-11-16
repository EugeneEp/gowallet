package controllers

import(
	"net/http"
	"app/utils"
	"app/models"
	"encoding/json"
)


// Отпарвить смс
func SendSms(w http.ResponseWriter, r *http.Request) {

	var sms models.Sms
	err := json.NewDecoder(r.Body).Decode(&sms)
	if err != nil {
		utils.Response(w, utils.ErrServerError)
	}
	newSms := sms.SendSms()
	utils.Response(w, newSms)

}

// Проверить код смс
func CheckSms(w http.ResponseWriter, r *http.Request) {
	var sms models.Sms
	err := json.NewDecoder(r.Body).Decode(&sms)
	if err != nil {
		utils.Response(w, utils.ErrServerError)
	}
	checkSms := sms.CheckSms()
	utils.Response(w, checkSms)
}
