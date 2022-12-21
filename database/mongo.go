package database

import (
	"context"
	"fmt"
	"log"

	userData "mateuszgua/to-do-list/database/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoMetaDataStore struct {
	Client         *mongo.Client
	DatabaseName   string
	CollectionName string
}

func NewMongoMetaDataStore(uri string, authSource string, authUserName string,
	authUserPassword string, databaseName string, collectionName string) (MongoMetaDataStore, error) {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri).SetAuth(
		options.Credential{
			AuthSource: authSource,
			Username:   authUserName,
			Password:   authUserPassword,
		}))
	if err != nil {
		return MongoMetaDataStore{}, fmt.Errorf("failed to connect mongodb client %w", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return MongoMetaDataStore{}, fmt.Errorf("failed to pinged mongo client %w", err)
	}

	store := MongoMetaDataStore{
		Client:         mongoClient,
		DatabaseName:   databaseName,
		CollectionName: collectionName,
	}
	if err != nil {
		return store, fmt.Errorf("failed to connect mongo meta data store %w", err)
	}

	log.Println("Successfully connected and pinged to mongodb.")
	return store, nil

}

func (store MongoMetaDataStore) SaveMetaData(metaData userData.UserMetaData) (string, error) {
	collection := store.Client.Database(store.DatabaseName).Collection(store.CollectionName)

	insertResult, err := collection.InsertOne(context.Background(), metaData)
	if err != nil {
		return "", fmt.Errorf("failed to save data in mongodb: %w", err)
	}

	idUser := insertResult.InsertedID.(primitive.ObjectID)
	return idUser.Hex(), nil
}

func (store MongoMetaDataStore) getUserMetaData(userId string) (userData.UserMetaData, error) {
	var mongoId primitive.ObjectID
	err := mongoId.UnmarshalText([]byte(userId))
	if err != nil {
		return userData.UserMetaData{}, fmt.Errorf("failed to unmarschal userID: %w", err)
	}

	collection := store.Client.Database(store.DatabaseName).Collection(store.CollectionName)
	filter := bson.D{{Key: "_id", Value: mongoId}}

	var result userData.UserMetaData
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("failed to get user data from database: %w", err)
	}

	return result, nil

}

func (store MongoMetaDataStore) GetUserMetaData(email string) (string, error) {
	currentUser, err := store.getUserMetaData(email)
	return currentUser.FirstName, err
}
