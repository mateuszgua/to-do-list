package user

import (
	"log"
	userData "mateuszgua/to-do-list/database/model"
	"mateuszgua/to-do-list/helpers"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// func userSignUp(mongoStore *mongo.MongoMetaDataStore, response http.ResponseWriter, request *http.Request) {
// response.Header().Set("Content-Type", "application/json")
// var user userData.UserMetaData
// json.NewDecoder(request.Body).Decode(&user)
// user.Password = getHash([]byte(user.Password))
// result, err := mongoStore.SaveMetaData(user)
// if err != nil {
// response.WriteHeader(http.StatusInternalServerError)
// response.Write([]byte(`{"message":"` + err.Error() + `"}`))
// return
// }
// json.NewEncoder(response).Encode(result)
// }
//
// func getHash(pwd []byte) string {
// hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
// if err != nil {
// log.Println("failed to generate hash password: %w", err)
//
// }
// return string(hash)
// }

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
		ID:       user.ID,
		Name:     user.Name,
		Nick:     user.Nick,
		Email:    user.Email,
		Accounts: accounts,
	}
	var token = prepareToken(user)
	var response = map[string]interface{}{"message": "all is fine"}
	response["jwt"] = token
	response["data"] = responseUser

	return response

}

func Register(name string, nick string, email string, pass string) map[string]interface{} {
	valid := helpers.Validation(
		[]userData.Validation{
			{Value: name, Valid: "name"},
			{Value: nick, Valid: "nick"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		generatedPassword := helpers.HashAndSalt([]byte(pass))
		user := &userData.UserMetaData{Name: name, Nick: nick, Email: email, Password: generatedPassword}

		account := &userData.Account{Type: "Daily account", Name: string(name + "'s" + "account"), Balance: 0, UserId: user.ID}

		accounts := []userData.ResponseAccount{}
		respAccount := userData.ResponseAccount{ID: account.UserId, Name: account.Name, Balance: int(account.Balance)}
		accounts = append(accounts, respAccount)

		var response = prepareResponse(user, accounts)

		return response
	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}
