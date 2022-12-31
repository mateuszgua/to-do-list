package helpers

import (
	"log"
	"regexp"

	userData "mateuszgua/to-do-list/database/model"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Println("failed to generate hash password: %w", err)
	}
	return string(hashed)
}

func Validation(values []userData.Validation) bool {
	firstName := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
	lastName := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$`)

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "firstname":
			if !firstName.MatchString(values[i].Value) {
				return false
			}
		case "lastname":
			if !lastName.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}

	return true
}
