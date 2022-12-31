package user

import (
	userData "mateuszgua/to-do-list/database/model"
	"mateuszgua/to-do-list/helpers"

	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY = []byte("gosecretkey")

// func userLogin(mongoStore *mongo.MongoMetaDataStore, response http.ResponseWriter, request *http.Request) {
// response.Header().Set("Content-Type", "application/json")
// var user userData.UserMetaData
// var dbUser userData.UserMetaData
//
// user := userData.UserMetaData{
// Email:    request.FormValue("email"),
// Password: request.FormValue("password"),
// }
// _ = user
// log.Println(request.FormValue("email"))
// log.Println(user.Password)
//
// json.NewDecoder(request.Body).Decode(&user)
// dbUser, err := mongoStore.GetUserMetaData(user.Email)
// if err != nil {
// response.WriteHeader(http.StatusInternalServerError)
// response.Write([]byte(`{"message":"` + err.Error() + `"}`))
// return
// }
//
// userPass := []byte(user.Password)
// dbPass := []byte(dbUser.Password)
// passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
// if passErr != nil {
// log.Println(passErr)
// response.Write([]byte(`{"response":"Wrong password!"}`))
// return
// }
//
// jwtToken, err := GenerateJWT()
// if err != nil {
// response.WriteHeader(http.StatusInternalServerError)
// response.Write([]byte(`{"message":"` + err.Error() + `"}`))
// return
// }
// response.Write([]byte(`{"token":"` + jwtToken + `"}`))
// }
//
// func GenerateJWT() (string, error) {
// token := jwt.New(jwt.SigningMethodHS256)
// tokenString, err := token.SignedString(SECRET_KEY)
// if err != nil {
// log.Println("Error in JWT token generation")
// return "", err
// }
// return tokenString, nil
// }

func Login(email string, pass string) map[string]interface{} {
	valid := helpers.Validation(
		[]userData.Validation{
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})

	if valid {
		user := &userData.UserMetaData{}
		//TODO: check if user exist in db
		// if db.Where("username = ? ", username).First(&user).RecordNotFound() {
		// return map[string]interface{}{"message": "User not found"}
		// }
		passError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if passError == bcrypt.ErrMismatchedHashAndPassword && passError != nil {
			return map[string]interface{}{"message": "Wrong password!"}
		}

		accounts := []userData.ResponseAccount{}
		// TODO find account for user
		// db.Table("accounts").Select("id, name, balance").Where("user_id = ? ", user.ID).Scan(&accounts)
		var response = prepareResponse(user, accounts)
		return response

	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}
