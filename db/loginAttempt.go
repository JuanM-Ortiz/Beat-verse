package db

import (
	"github.com/JuanM-Ortiz/beat-verse/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginAttempt(email string, password string) (models.User, bool) {
	user, found, _ := CheckUser(email)

	if !found {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
