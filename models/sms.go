package models

import (
	"database/sql"
	"app/utils"
	"time"
	re "regexp"
)

type Sms struct{
	id int
	UserId int
	Code int
	Action string
	Time int64
	Status int
	Phone string
}

func(s *Sms)ExtractDigits() {
	var replace = re.MustCompile("[^0-9]")
    s.Phone = replace.ReplaceAllString(s.Phone, "")
}

func(s *Sms)SendSms() map[string]interface{}{
	currentTime := time.Now().Unix()
	s.ExtractDigits()
	var id int
	err := DB.QueryRow(`SELECT id FROM public.sms_approve WHERE phone=$1 AND action=$2`, s.Phone, s.Action).Scan(&id)
	if err != sql.ErrNoRows{
		if (currentTime - s.Time) < 60 {
			return utils.ErrSmsTimelimit
		}else if (s.Action == "reg" && s.Status == 1) {
			return utils.ErrSmsApprove
		}else{
			sqlStatementUpdate := `UPDATE public.sms_approve SET time=$1 WHERE id=$2`
			_, errUpdate := DB.Exec(sqlStatementUpdate, currentTime, id)
			if errUpdate != nil{
				return utils.ErrServerError
			}
			return map[string]interface{}{"success":true,"msg":"Смс успешно отправлено"}
		}
	}
	if err != nil && err != sql.ErrNoRows {
        return utils.ErrServerError
    }
    sqlStatement := `INSERT INTO public.sms_approve (user_id, code, action, time, status, phone) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`
	errInsert := DB.QueryRow(sqlStatement, nil, 1111, s.Action, currentTime, 0, s.Phone).Scan(&id)
	if errInsert != nil{
		return utils.ErrServerError
	}
    return map[string]interface{}{"success":true,"msg":"Смс успешно отправлено"}
}

func(s *Sms)CheckSms() map[string]interface{}{
	currentTime := time.Now().Unix()
	s.ExtractDigits()
	var id int
	err := DB.QueryRow(`SELECT id, time FROM public.sms_approve WHERE status=0 AND phone=$1 AND action=$2 AND code=$3`, s.Phone, s.Action, s.Code).Scan(&id, &s.Time)
	if err == sql.ErrNoRows{
		return utils.ErrSmsCode
	}
	if err != nil {
		return utils.ErrServerError
	}
	if (currentTime - s.Time) > (60 * 60) {
		return utils.ErrSmsExpired
	}
	sqlStatementUpdate := `UPDATE public.sms_approve SET status=1 WHERE id = $1`
	_, errUpdate := DB.Exec(sqlStatementUpdate, id)
	if errUpdate != nil{
		return utils.ErrServerError
	}
	return map[string]interface{}{"success":true,"msg":"Код смс подтвержден"}
}