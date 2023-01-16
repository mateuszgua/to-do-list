package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"mateuszgua/to-do-list/handle"
)

var router = mux.NewRouter()

func MyRouter() (*mux.Router, error) {

	// router.HandleFunc("/api/hello", helloHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/", handle.IndexPageHandler).Methods("GET")
	router.HandleFunc("/index", handle.IndexPageHandler).Methods("GET")
	router.HandleFunc("/panel", handle.PanelPageHandler).Methods("GET")
	router.HandleFunc("/panel", handle.PanelHandler).Methods("POST")
	router.HandleFunc("/register", handle.RegisterPageHandler).Methods("GET")
	router.HandleFunc("/register", handle.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handle.LoginPageHandler).Methods("GET")
	router.HandleFunc("/login", handle.LoginHandler).Methods("POST")
	// router.HandleFunc("/api/user/task", helloHandler).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/user/task", helloHandler).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/user/task/{id}", helloHandler).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/user/undoTask/{id}", helloHandler).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/user/deleteTask/{id}", helloHandler).Methods("DELETE", "OPTIONS")
	// router.HandleFunc("/api/user/deleteAllTask", helloHandler).Methods("DELETE", "OPTIONS")

	http.Handle("/", router)
	return router, nil
}
