package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	mongodb "mateuszgua/to-do-list/database"
	userData "mateuszgua/to-do-list/database/model"

	"github.com/joho/godotenv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	godotenv.Load(".env")

	userName := os.Getenv("MONGO_USERNAME")
	userPassword := os.Getenv("MONGO_PASSWORD")
	databaseName := os.Getenv("MONGO_DATABASE_NAME")
	collectionName := os.Getenv("MONGO_COLLECTION_NAME")
	mongoHost := os.Getenv("MONGO_HOST")
	mongoServerPort := os.Getenv("MONGO_PORT")
	authSource := os.Getenv("MONGO_AUTH_SOURCE")
	authUserName := os.Getenv("MONGO_AUTH_USER")
	authUserPassword := os.Getenv("MONGO_AUTH_PASSWORD")

	// add collection name for user and for task
	uri := fmt.Sprintf("mongodb://%s:%s@%s%s", userName, userPassword, mongoHost, mongoServerPort)

	mongoStore, err := mongodb.NewMongoMetaDataStore(uri, authSource, authUserName, authUserPassword, databaseName, collectionName)
	if err != nil {
		log.Fatal("failed to create new mongo client", err)
	}

	currentTime := time.Now()

	testSaveDataInDb := userData.UserMetaData{
		FirstName:      "Jan",
		LastName:       "Kowalski",
		Password:       "Qwerty1234",
		Email:          "jankowalski@gmail.com",
		IndexationDate: currentTime,
	}

	mongoUserId, _ := mongoStore.SaveMetaData(testSaveDataInDb)
	log.Println(mongoUserId)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	serverUrl := fmt.Sprintf(":%s", httpPort)
	log.Printf("Starting server on http://localhost%s", serverUrl)

	err = http.ListenAndServe(serverUrl, nil)
	if err != nil {
		log.Fatal("failed to start server", err)
	}

}
