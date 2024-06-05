package db

import (
	"context"
	"fmt"

	"github.com/MikeB1124/registration-lambda/configuration"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
	config := configuration.GetConfig()
	// Connect to MongoDB
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.du0vf.mongodb.net", config.MongoDB.Username, config.MongoDB.Password))
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	mongoClient = client
}

func InsertNewUser(user User) error {
	collection := mongoClient.Database("registrationDB").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func ValidLogin(user User) (bool, error) {
	collection := mongoClient.Database("registrationDB").Collection("users")
	filter := bson.M{"username": user.Username, "password": user.Password}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func UserExists(user User) (bool, error) {
	collection := mongoClient.Database("registrationDB").Collection("users")
	filter := bson.M{"username": user.Username}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func UpdateLastLogin(lastLoginTime string, username string) error {
	collection := mongoClient.Database("registrationDB").Collection("users")
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"lastlogin": lastLoginTime}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}
