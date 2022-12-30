package user

import (
	"encoding/json"
	"log"
	mongo "mateuszgua/to-do-list/database"
	userData "mateuszgua/to-do-list/database/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func userSignUp(mongoStore *mongo.MongoMetaDataStore, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user userData.UserMetaData
	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))
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

func prepareToken(user *userData.UserMetaData) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	if err != nil {
		log.Println("failed to get token password: %w", err)
	}
	return token
}

func prepareResponse(user *userData.UserMetaData, accounts []userData.ResponseAccount) map[string]interface{} {
	responseUser := *&userData.ResponseUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Accounts:  accounts,
	}
	var token = prepareToken(user)
	var response = map[string]interface{}{"message": "all is fine"}
	response["jwt"] = token
	response["data"] = responseUser

	return response

}

// func Register(firstName string, lastName string, email string, pass string) map[string]interface{} {
// valid :=
// }
