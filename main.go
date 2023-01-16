package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mongodb "mateuszgua/to-do-list/database"
	userData "mateuszgua/to-do-list/database/model"
	"mateuszgua/to-do-list/server"
	"mateuszgua/to-do-list/server/router"

	"github.com/joho/godotenv"
)

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

	uri := fmt.Sprintf("mongodb://%s:%s@%s%s", userName, userPassword, mongoHost, mongoServerPort)

	httpPort := os.Getenv("HTTP_PORT")

	router, err := router.MyRouter()
	if err != nil {
		log.Fatal("failed to add router", err)
	}

	err = server.MyServer(httpPort, router)
	if err != nil {
		log.Fatal("failed connect with server", err)
	}

	mongoStore, err := mongodb.NewMongoMetaDataStore(uri, authSource, authUserName, authUserPassword, databaseName, collectionName)
	if err != nil {
		log.Fatal("failed to create new mongo client", err)
	}

	// _ = mongoStore
	currentTime := time.Now()

	testSaveDataInDb := userData.UserMetaData{
		Name:           "Jan",
		Nick:           "Janko",
		Password:       "Qwerty1234",
		Email:          "jankowalski@gmail.com",
		IndexationDate: currentTime,
	}

	mongoUserId, _ := mongoStore.SaveMetaData(testSaveDataInDb)
	log.Println(mongoUserId)

}
