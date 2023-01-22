package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/anujshah3/AddressTrail/config"
	"github.com/anujshah3/AddressTrail/models"
)

func AddUser(user *models.User) error {

	client, err := config.GetMongoDBClient()
	if err != nil {
		return err
	}

	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")

	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}


func AddNewAddressToUser(userID string, address *models.AddressWithDates) error {

	client, err := config.GetMongoDBClient()
	if err != nil {
		return err
	}

	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")

	filter := bson.M{"id": userID}
	update := bson.M{"$push": bson.M{"addresses": address}}
	_, err = userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
