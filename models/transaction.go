package models

import (
	"app/utils"
	"database/sql"
	"strconv"
	"time"
	"crypto/md5"
	"encoding/hex"
	"encoding/csv"
	"os"
	"log"
)

type Transaction struct{
	id int
	WalletSender int
	WalletReciever int
	UserId int
	MovementType string
	Amount string
	Comission string
	Time int64
	secretToken string
	Status int
	email string
}

func DateToTime(str string) int64 {
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)
    if err != nil {
        return 0
    }
    return t.Unix()
}

func CreateCsv(filename string, rows [][]string){
	csvFile, errCsv := os.Create("static/csv/"+filename)
	if errCsv != nil{
		log.Fatal("Failed to create file")
	}
	csvWriter := csv.NewWriter(csvFile)
	errCsvWrite := csvWriter.WriteAll(rows)
	if errCsvWrite != nil{
		log.Fatal("Failed to write csv")
	}
	csvWriter.Flush()
	csvFile.Close()
}

func(t *Transaction)HashCsv()string{
	hash := md5.Sum([]byte(string(t.UserId) + "IPOjk0921j01293j120j"))
	filename := hex.EncodeToString(hash[:])
	return filename + ".csv"
}

func(t *Transaction)GetAll(filters map[string][]string)map[string]interface{}{
	sqlStatement := `SELECT id, movement_type, amount, time, status FROM public.transactions WHERE user_id=$1`
	date_from, okFrom := filters["date_from"]
	if okFrom || len(date_from) == 1{
		date_from := DateToTime(date_from[0])
		sqlStatement = sqlStatement + ` AND time > ` + strconv.FormatInt(date_from, 10)
	}
	date_to, okTo := filters["date_to"]
	if okTo || len(date_to) == 1{
		date_to := DateToTime(date_to[0])
		sqlStatement = sqlStatement + ` AND time < ` + strconv.FormatInt(date_to, 10)
	}
	sqlStatement = sqlStatement + ` ORDER BY id `
	_, okAsc := filters["asc"]
	if okAsc{
		sqlStatement = sqlStatement + ` ASC `
	}else{
		sqlStatement = sqlStatement + ` DESC `
	}
	size, okSize := filters["size"]
	if okSize || len(size) == 1{
		sqlStatement = sqlStatement + ` LIMIT ` + size[0]
	}
	result, err := DB.Query(sqlStatement, t.UserId)
	if err == sql.ErrNoRows{
		return map[string]interface{}{"success":true,"transactions":nil}
	}
	if err != nil{
		return utils.ErrServerError
	}
	arr := map[string]interface{}{}

	for result.Next(){
		var tr Transaction
		err := result.Scan(&tr.id, &tr.MovementType, &tr.Amount, &tr.Time, &tr.Status)
		if err != nil {
			return utils.ErrServerError
		}
		getTr := map[string]interface{}{
			"id": tr.id,
			"movement_type": tr.MovementType,
			"amount": tr.Amount,
			"time": time.Unix(tr.Time, 0),
			"status": tr.Status,
		}

		arr[strconv.Itoa(tr.id)] = getTr
	}
	return map[string]interface{}{"success":true,"transactions":arr}
}

func(t *Transaction)GetAllCsv(filters map[string][]string)map[string]interface{}{
	sqlStatement := `SELECT id, movement_type, amount, time, status FROM public.transactions WHERE user_id=$1`
	date_from, okFrom := filters["date_from"]
	if okFrom || len(date_from) == 1{
		date_from := DateToTime(date_from[0])
		sqlStatement = sqlStatement + ` AND time > ` + strconv.FormatInt(date_from, 10)
	}
	date_to, okTo := filters["date_to"]
	if okTo || len(date_to) == 1{
		date_to := DateToTime(date_to[0])
		sqlStatement = sqlStatement + ` AND time < ` + strconv.FormatInt(date_to, 10)
	}
	sqlStatement = sqlStatement + ` ORDER BY id `
	_, okAsc := filters["asc"]
	if okAsc{
		sqlStatement = sqlStatement + ` ASC `
	}else{
		sqlStatement = sqlStatement + ` DESC `
	}
	size, okSize := filters["size"]
	if okSize || len(size) == 1{
		sqlStatement = sqlStatement + ` LIMIT ` + size[0]
	}
	result, err := DB.Query(sqlStatement, t.UserId)
	if err == sql.ErrNoRows{
		return map[string]interface{}{"success":true,"transactions":nil}
	}
	if err != nil{
		return utils.ErrServerError
	}
	rows := [][]string{
		{"Id", "Тип тарнзакции", "Сумма", "Дата", "Статус"},
	}

	for result.Next(){
		var tr Transaction
		err := result.Scan(&tr.id, &tr.MovementType, &tr.Amount, &tr.Time, &tr.Status)
		if err != nil {
			return utils.ErrCSVError
		}

		rows = append(rows, []string{
			strconv.Itoa(tr.id), 
			tr.MovementType, 
			tr.Amount, 
			time.Unix(tr.Time, 0).String(), 
			strconv.Itoa(tr.Status),
		})
	}

	filename := t.HashCsv()
	go CreateCsv(filename, rows)

	return map[string]interface{}{
		"success":true,
		"file": "http://localhost:8080/static/csv/" + filename,
	}
}

func(t *Transaction)GetOne(id string)map[string]interface{}{
	tid, errInt := strconv.Atoi(id)
	if errInt != nil{
		return utils.ErrTransactionId
	}
	sqlStatement := `SELECT movement_type, amount, time, status FROM public.transactions WHERE user_id=$1 AND id=$2`
	err := DB.QueryRow(sqlStatement, t.UserId, tid).Scan(&t.MovementType, &t.Amount, &t.Time, &t.Status)
	if err == sql.ErrNoRows{
		return map[string]interface{}{"success":true,"transaction":nil}
	}
	if err != nil{
		return utils.ErrServerError
	}
	arr := map[string]interface{}{
		"id":tid,
		"movement_type":t.MovementType,
		"amount":t.Amount,
		"time":time.Unix(t.Time, 0),
		"status":t.Status,
	}

	return map[string]interface{}{"success":true,"transaction":arr}
}