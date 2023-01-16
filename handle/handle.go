package handle

import (
	"fmt"
	"net/http"

	helpers "mateuszgua/to-do-list/helpers"
)

type Login struct {
	Email    string
	Password string
}

type Register struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type ErrResponse struct {
	Message string
}

func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	body, _ := helpers.LoadFile("templates/index.html")
	fmt.Fprintf(response, body)
}

func PanelPageHandler(response http.ResponseWriter, request *http.Request) {
	body, _ := helpers.LoadFile("templates/panel.html")
	fmt.Fprintf(response, body)
}

func PanelHandler(response http.ResponseWriter, request *http.Request) {
	title := request.FormValue("title")
	description := request.FormValue("description")
	startDate := request.FormValue("startDate")
	endDate := request.FormValue("endDate")
	priority := request.FormValue("priority")
	status := request.FormValue("status")

	_title := false
	_description := false
	_startDate := false
	_endDate := false
	_priority := false
	_status := false

	_title = !helpers.IsEmpty(title)
	_description = !helpers.IsEmpty(description)
	_startDate = !helpers.IsEmpty(startDate)
	_endDate = !helpers.IsEmpty(endDate)
	_priority = !helpers.IsEmpty(priority)
	_status = !helpers.IsEmpty(status)

	if _title && _description && _startDate && _endDate && _priority && _status {
		fmt.Fprintln(response, "Title: ", title)
		fmt.Fprintln(response, "Description: ", description)
		fmt.Fprintln(response, "Start date: ", startDate)
		fmt.Fprintln(response, "End date: ", endDate)
		fmt.Fprintln(response, "Priority: ", priority)
		fmt.Fprintln(response, "Status: ", status)
	} else {
		fmt.Fprintln(response, "This fields can not be blank! Please fill all fields.")

	}
}

func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	body, _ := helpers.LoadFile("templates/login.html")
	fmt.Fprintf(response, body)
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if !helpers.IsEmpty(email) && !helpers.IsEmpty(pass) {
		_userIsValid := helpers.UserIsValid(email, pass)

		if _userIsValid {
			redirectTarget = "/panel"
		} else {
			redirectTarget = "/register"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
	body, _ := helpers.LoadFile("templates/register.html")
	fmt.Fprintf(response, body)
}

func RegisterHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	userName := request.FormValue("name")
	userNick := request.FormValue("nick")
	email := request.FormValue("email")
	pwd := request.FormValue("password")
	confirmPwd := request.FormValue("confirmPassword")

	_userName := false
	_userNick := false
	_email := false
	_pwd := false
	_confirmPwd := false

	_userName = !helpers.IsEmpty(userName)
	_userNick = !helpers.IsEmpty(userNick)
	_email = !helpers.IsEmpty(email)
	_pwd = !helpers.IsEmpty(pwd)
	_confirmPwd = !helpers.IsEmpty(confirmPwd)

	if _userName && _userNick && _email && _pwd && _confirmPwd {
		fmt.Fprintln(response, "Username: ", userName)
		fmt.Fprintln(response, "Nick: ", userNick)
		fmt.Fprintln(response, "Email: ", email)
		fmt.Fprintln(response, "Password: ", pwd)
		fmt.Fprintln(response, "ConfirmPassword: ", confirmPwd)
	} else {
		fmt.Fprintln(response, "This fields can not be blank!")
	}
}

// func UserLogin(w http.ResponseWriter, r *http.Request) {
// 	// body, err := ioutil.ReadAll(r.Body)
// 	body, err := ioutil.ReadFile("templates/login.html")
// 	if err != nil {
// 		log.Println("failed to read data from login form: %w", err)
// 	}

// 	var formattedBody Login
// 	err = json.Unmarshal(body, &formattedBody)
// 	if err != nil {
// 		log.Println("failed to read data from login json: %w", err)
// 	}
// 	login := users.Login(formattedBody.Email, formattedBody.Password)

// 	if login["message"] == "all is fine" {
// 		resp := login
// 		json.NewEncoder(w).Encode(resp)
// 	} else {
// 		resp := ErrResponse{Message: "Wrong username or password"}
// 		json.NewEncoder(w).Encode(resp)
// 	}
// }

// func UserRegister(w http.ResponseWriter, r *http.Request) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Println("failed to read data from register form: %w", err)
// 	}

// 	var formattedBody Register
// 	err = json.Unmarshal(body, &formattedBody)
// 	if err != nil {
// 		log.Println("failed to read data from register json: %w", err)
// 	}

// 	register := users.Register(formattedBody.FirstName, formattedBody.LastName, formattedBody.Email, formattedBody.Password)

// 	if register["message"] == "all is fine" {
// 		resp := register
// 		json.NewEncoder(w).Encode(resp)
// 	} else {
// 		resp := ErrResponse{Message: "Wrong username or password"}
// 		json.NewEncoder(w).Encode(resp)
// 	}

// }
