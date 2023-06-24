package services

import (
	"context"

	"github.com/anujshah3/AddressTrail/config"
	"github.com/anujshah3/AddressTrail/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertAddress(address *models.Address) error {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	addressCollection := config.GetCollection(client, "address")

	_, err = addressCollection.InsertOne(context.Background(), address)
	if err != nil {
		return err
	}

	return nil
}


func DeleteAddress(addressID string) error {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	addressCollection := config.GetCollection(client, "address")

	filter := bson.M{"id": addressID}
	_, err = addressCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}