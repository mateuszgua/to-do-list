package router

import (
	"fmt"
	"net/http"

	"mateuszgua/to-do-list/handle"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func MyRouter(httpPort string) (*mux.Router, error) {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	router := mux.NewRouter()

	router.HandleFunc("/api/hello", helloHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/login", handle.UserLogin)
	router.HandleFunc("/api/user/register", helloHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/task", helloHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/task", helloHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/task/{id}", helloHandler).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/user/undoTask/{id}", helloHandler).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/user/deleteTask/{id}", helloHandler).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/user/deleteAllTask", helloHandler).Methods("DELETE", "OPTIONS")

	return router, nil
}
