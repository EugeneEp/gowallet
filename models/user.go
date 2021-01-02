package models

import(
	"database/sql"
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"app/utils"
	"time"
	re "regexp"
	"strings"
	"encoding/json"
	"net/http"
	"io"
	"os"
	"strconv"
	"fmt"
)

type User struct{
	id int
	Email string
	Phone string
	Password string
	Confirm string
	Identity struct{
		Fullname string `json:"fullname"`
		Passport int `json:"passport"`
		PassportIssuedAt string `json:"passportIssuedAt"`
	} `json:"identity"`
	roots int
	Token string
}

func(u *User)UserExistByPhone() (int, error){
	var errRDB error
	id, errRDB := RDB.Get(u.Phone).Result()
	if errRDB != nil{
		fmt.Println("no redis")
		errDB := DB.QueryRow(`SELECT id FROM public.users WHERE phone=$1`, u.Phone).Scan(&id)
		if errDB == sql.ErrNoRows{
			return 0, errDB
		}
		if errDB != nil {
	        return 0, errDB
	    }
	    _ = RDB.Set(u.Phone, id, 0).Err()
	}else{
		fmt.Println("redis")
	}
	result, _ := strconv.Atoi(id)
    return result, nil
}

func(u *User)CreateToken()(string, error){
	var err error
	secret := os.Getenv("SECRET_TOKEN")

	atClaims := jwt.MapClaims{}
	atClaims["phone"] = u.Phone
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func(u *User)Init() {
	hash := md5.Sum([]byte(u.Password + os.Getenv("SALT")))
	u.Password = hex.EncodeToString(hash[:])
	if u.Confirm != ""{
		hashConfirm := md5.Sum([]byte(u.Confirm + os.Getenv("SALT")))
		u.Confirm = hex.EncodeToString(hashConfirm[:])
	}
	var replace = re.MustCompile("[^0-9]")
    u.Phone = replace.ReplaceAllString(u.Phone, "")
}

func(u *User)Login()map[string]interface{} {
	var id int
	err := DB.QueryRow(`SELECT id FROM public.users WHERE phone=$1 AND password=$2`, u.Phone, u.Password).Scan(&id)
	if err == sql.ErrNoRows{
		return utils.ErrIncorrectLogin
	}
	if err != nil {
        return utils.ErrServerError
    }
    u.Token, err = u.CreateToken()
    if err != nil {
        return utils.ErrCreateToken
    }
    return map[string]interface{}{"success":true,"phone":u.Phone, "token":u.Token}
}

func(u *User)CreateUser()map[string]interface{}{
	if u.Password != u.Confirm{
		return utils.ErrPassMatch
	}
	_, errExist := u.UserExistByPhone()
	if errExist != sql.ErrNoRows{
		return utils.ErrRegExist
	}
	var sms Sms
	sqlStatementSms := `SELECT id, status FROM public.sms_approve WHERE phone=$1 AND action='reg' AND status=1`
	errSms := DB.QueryRow(sqlStatementSms, u.Phone).Scan(&sms.id, &sms.Status)
	if errSms == sql.ErrNoRows{
		return utils.ErrSmsConfirm
	}
	if errSms != nil {
        return utils.ErrServerError
    }
	var id int
	sqlStatement := `INSERT INTO public.users (phone, password, confirmation, roots) VALUES ($1, $2, $3, $4) RETURNING id`
	err := DB.QueryRow(sqlStatement, u.Phone, u.Password, 0, 1).Scan(&id)
	if err != nil {
        return utils.ErrServerError
    }
    u.Token, err = u.CreateToken()
    if err != nil {
        return utils.ErrCreateToken
    }
    return map[string]interface{}{"success":true,"email":u.Phone, "token":u.Token}
}

func(u *User)UpdateUser(id int)map[string]interface{}{
	if u.Identity.Fullname != ""{
		explode := strings.Fields(u.Identity.Fullname)
		if len(explode) < 3{
			return utils.ErrFullname
		}
	}
	identity, err := json.Marshal(u.Identity)
	if err != nil{
		return utils.ErrServerError
	}
	sqlStatement := `UPDATE users SET identity=$1 WHERE id=$2`
	_, errUpdate := DB.Exec(sqlStatement, identity, id)
	if errUpdate != nil{
		return utils.ErrServerError
	}
	return map[string]interface{}{"success":true,"msg":"Данные успешно обновлены"}
}

func(u *User)ProfilePicture(r *http.Request, user_id int)map[string]interface{}{
	file, _, err := r.FormFile("img")
    fileName := "static/profile/img/" + strconv.Itoa(user_id) + ".png"
    if err != nil {
        return utils.ErrServerError
    }
    defer file.Close()

    f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        return utils.ErrServerError
    }
    defer f.Close()
    _, _ = io.Copy(f, file)
    return map[string]interface{}{"success":true,"msg":"Файл успешно загружен"}
}