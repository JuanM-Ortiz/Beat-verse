package routers

import (
	"errors"
	"strings"

	"github.com/JuanM-Ortiz/beat-verse/db"
	"github.com/JuanM-Ortiz/beat-verse/models"
	"github.com/dgrijalva/jwt-go"
)

var Email string

var UserID string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myPassword := []byte("JuanManuel_92")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myPassword, nil
	})

	if err == nil {
		_, found, _ := db.CheckUser(claims.Email)
		if found {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
