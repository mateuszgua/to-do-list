package handle

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	users "mateuszgua/to-do-list/account"

	"github.com/gorilla/sessions"
)

type Register struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type ErrResponse struct {
	Message string
}

var store = sessions.NewCookieStore([]byte("mysession"))

func UserLogin(w http.ResponseWriter, r *http.Request) {
	// tmp, err := template.ParseFiles("views/login.html")
	// if err != nil {
	// log.Printf("failed to get login template: %s", err)
	// }
	// tmp.Execute(w, nil)
	tmpl := template.Must(template.ParseFiles("views/login.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	// r.ParseForm()

	// user := userData.UserMetaData{
	// 	Email:    r.FormValue("email"),
	// 	Password: r.FormValue("password"),
	// }
	// _ = user
	// log.Println(r.FormValue("email"))
	// log.Println(user.Password)

	tmpl.Execute(w, struct{ Success bool }{true})
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("failed to read data from form: %w", err)
	}

	var formattedBody Register
	err = json.Unmarshal(body, &formattedBody)
	if err != nil {
		log.Println("failed to read data from json: %w", err)
	}

	register := users.Register(formattedBody.FirstName, formattedBody.LastName, formattedBody.Email, formattedBody.Password)

	if register["message"] == "all is fine" {
		resp := register
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}

}
