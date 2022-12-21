package user

import (
	"encoding/json"
	"log"
	mongo "mateuszgua/to-do-list/database"
	userData "mateuszgua/to-do-list/database/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY = []byte("gosecretkey")

func userLogin(mongoStore *mongo.MongoMetaDataStore, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user userData.UserMetaData
	var dbUser userData.UserMetaData
	json.NewDecoder(request.Body).Decode(&user)
	dbUser, err := mongoStore.GetUserMetaData(user.Email)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)
	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	if passErr != nil {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong password!"}`))
		return
	}

	jwtToken, err := GenerateJWT()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}
