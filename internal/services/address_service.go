package services

import (
	"context"

	"github.com/anujshah3/AddressTrail/config"
	"github.com/anujshah3/AddressTrail/internal/models"
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


