package user

import (
	"encoding/json"
	"log"
	mongo "mateuszgua/to-do-list/database"
	userData "mateuszgua/to-do-list/database/model"
	"net/http"

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
