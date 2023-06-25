package services

import (
	"context"

	"github.com/anujshah3/AddressTrail/config"
	"github.com/anujshah3/AddressTrail/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertAddress(address *models.Address) (string, error) {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return "", err
	}
	defer client.Disconnect(context.Background())

	addressCollection := config.GetCollection(client, "address")

	result, err := addressCollection.InsertOne(context.Background(), address)
	if err != nil {
		return "", err
	}

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	return insertedID, nil
}


func DeleteAddress(addressID string) error {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	addressCollection := config.GetCollection(client, "address")

	objID, err := primitive.ObjectIDFromHex(addressID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	_, err = addressCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func GetAddressByID(addressID string) (*models.Address, error) {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	addressCollection := config.GetCollection(client, "address")

	objID, err := primitive.ObjectIDFromHex(addressID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	result := addressCollection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var address models.Address
	err = result.Decode(&address)
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func UpdateAddress(addressID string, address *models.Address) error {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	addressCollection := config.GetCollection(client, "address")
	objID, err := primitive.ObjectIDFromHex(addressID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": address}
	_, err = addressCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
