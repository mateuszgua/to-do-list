package handle

import (
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

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
