package repository

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

func (r *TopGamesPostgres) Login(n string,p string) (string,error) {
	var truepass string
	res, err := r.db.Query("select * from users where Name = $1", n)
	defer res.Close()
	if err != nil {
		return "", err
	}
	for res.Next() {
		err = res.Scan(&n,&truepass)
		if err != nil {
			return "", err
		}
	}
	if p==truepass{
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = n
		claims["exp"] = time.Now().Add(time.Hour).Unix()
		t, err := token.SignedString([]byte("sirok"))
		if err != nil {
			return "",err
		}
		return t,nil
	}
	return "",echo.ErrUnauthorized
}