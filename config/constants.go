package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetEnvVal(EnvVar string) string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv(EnvVar)
}

func SetupConfig() *oauth2.Config{

	ClientID := GetEnvVal("GOOGLE_CLIENT_ID")	
	ClientSecret := GetEnvVal("GOOGLE_CLIENT_SECRET")
	RedirectURL := GetEnvVal("AUTH_CALLBACK_URL")

	conf := &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}

func GetMongoDBClient() (*mongo.Client, error) {
	dbConfigURI := GetEnvVal("MONGO_URI")

	clientOptions := options.Client().ApplyURI(dbConfigURI)

	ctx := context.Background()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}


func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("address_trail").Collection(collectionName)
    return collection
}