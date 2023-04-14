package jwt

import (
	"time"

	"github.com/JuanM-Ortiz/beat-verse/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJwt(t models.User) (string, error) {

	myPassword := []byte("JuanManuel_92")

	payload := jwt.MapClaims{
		"email":       t.Email,
		"name":        t.Name,
		"lastName":    t.LastName,
		"dateOfBirth": t.DateOfBirth,
		"biography":   t.Biography,
		"location":    t.Location,
		"webSite":     t.WebSite,
		"_id":         t.ID.Hex(),
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myPassword)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
