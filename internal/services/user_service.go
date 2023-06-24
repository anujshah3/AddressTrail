package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/anujshah3/AddressTrail/config"
	"github.com/anujshah3/AddressTrail/internal/models"
)


func AddUser(user *models.User) (string, error) {

	client, err := config.GetMongoDBClient()
	if err != nil {
		return "", err
	}

	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")
	existingUser := models.User{}
	err = userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		res, err := userCollection.InsertOne(context.Background(), user)
		if err != nil {
			return "", err
		}
		userID := res.InsertedID.(primitive.ObjectID).Hex()
		return userID, nil
	}
	return existingUser.ID, nil
}


func DeleteUser(email string) error {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return err
	}

	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")

	existingUser := models.User{}
	err = userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		return fmt.Errorf("user with email '%s' does not exist", email)
	} else if err != nil {
		return err
	}

	_, err = userCollection.DeleteOne(context.Background(), bson.M{"email": email})
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


func DeleteAddressFromUser(userID string, address *models.AddressWithDates) error {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return err
	}

	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")

	filter := bson.M{"id": userID}
	update := bson.M{"$pull": bson.M{"addresses": address}}
	_, err = userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}


func GetUserAddresses(email string) ([]*models.AddressWithDates, error) {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")

	filter := bson.M{"email": email}
	projection := bson.M{"addresses": 1}

	result := []*models.AddressWithDates{}
	cursor, err := userCollection.Find(context.Background(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user struct {
			Addresses []*models.AddressWithDates `bson:"addresses"`
		}
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		result = append(result, user.Addresses...)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
