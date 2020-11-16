package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"app/utils"
	"strings"
	"context"
	"app/models"
	"time"
)

type Token struct {
	Phone string
	Exp int64
	jwt.StandardClaims
}

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


		notAuth := []string{"/auth/login", "/auth/registration", "/sms/send", "/sms/check", "/static/csv"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if strings.Contains(requestPath, value) {
				next.ServeHTTP(w, r)
				return
			}
		}


		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			utils.Response(w, utils.ErrTokenMissing)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			utils.Response(w, utils.ErrSendToken)
			return
		}
		authType := splitted[0]
		if authType != "Bearer"{
			utils.Response(w, utils.ErrAuthType)
			return
		}

		secret := "Jawe21321dawdawd="

		tokenPart := splitted[1]
		tk := &Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			utils.Response(w, utils.ErrMalformedToken)
			return
		}

		if !token.Valid {
			utils.Response(w, utils.ErrInvalidToken)
			return
		}

		if tk.ExpiresAt < time.Now().Unix() {
			utils.Response(w, utils.ErrExpiredToken)
			return
		}

		var user = models.User{Phone:tk.Phone}
		id, user_err := user.UserExistByPhone()
		if user_err != nil{
			utils.Response(w, utils.ErrUserNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	});
}
