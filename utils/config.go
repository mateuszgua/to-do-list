package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UserName         string
	UserPassword     string
	DatabaseName     string
	CollectionName   string
	MongoHost        string
	MongoServerPort  string
	AuthSource       string
	AuthUserName     string
	AuthUserPassword string
	Uri              string
	HttpPort         string
}

func LoadConfig() Config {

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
	httpPort := os.Getenv("HTTP_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s@%s%s", userName, userPassword, mongoHost, mongoServerPort)

	config := Config{
		UserName:         userName,
		UserPassword:     userPassword,
		DatabaseName:     databaseName,
		CollectionName:   collectionName,
		MongoHost:        mongoHost,
		MongoServerPort:  mongoServerPort,
		AuthSource:       authSource,
		AuthUserName:     authUserName,
		AuthUserPassword: authUserPassword,
		Uri:              uri,
		HttpPort:         httpPort,
	}

	return config
}
