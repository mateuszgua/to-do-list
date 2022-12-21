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

func userSignUp(mongoStore *mongo.MongoMetaDataStore, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user userData.UserMetaData
	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := mongoStore.SaveMetaData(user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println("failed to generate hash password: %w", err)

	}
	return string(hash)
}

var SECRET_KEY = []byte("gosecretkey")

func userLogin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user userData.UserMetaData
	var dbUser userData.UserMetaData
	json.NewDecoder(request.Body).Decode(&user)
	user, err := mongo.MongoMetaDataStore.GetUserMetaData()
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
