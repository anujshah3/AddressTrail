package services

import (
	"context"
	"time"

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
		opts := options.InsertOne()
		res, err := userCollection.InsertOne(context.Background(), user, opts)
		if err != nil {
			return "", err
		}
		userID := res.InsertedID.(primitive.ObjectID).Hex()
		return userID, nil
	}
	return existingUser.ID, nil
}


func DeleteUser(userID string) error {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return err
	}

	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = userCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
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

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
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
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$pull": bson.M{"addresses": address}}
	_, err = userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}



func GetUserAddresses(userID string) ([]*models.UserAddresses, error) {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	// projection := bson.M{"addresses": 1}
	projection := bson.M{"addresses.addressid": 1, "addresses.startdate": 1, "addresses.enddate": 1}

	result := []*models.UserAddresses{}
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

		for _, address := range user.Addresses {
			addressDetails, err := GetAddressByID(address.AddressID)
			if err != nil {
				return nil, err
			}

			userAddress := &models.UserAddresses{
				AddressID:  address.AddressID,
				Street:     addressDetails.Street,
				Unit:       addressDetails.Unit,
				City:       addressDetails.City,
				State:      addressDetails.State,
				PostalCode: addressDetails.PostalCode,
				Country:    addressDetails.Country,
				StartDate:  address.StartDate,
				EndDate:    address.EndDate,
			}
			result = append(result, userAddress)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}


func UpdateFilteredAddresses(userID string, addressID string, startDate time.Time, endDate time.Time, newStartDate time.Time, newEndDate time.Time) error {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")

	filter := bson.M{
		"_id": userID,
		"addresses": bson.M{
			"$elemMatch": bson.M{
				"addressID": addressID,
				"startDate": startDate,
				"endDate":   endDate,
			},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"addresses.$.startDate": newStartDate,
			"addresses.$.endDate":   newEndDate,
			"addresses.$.updatedAt": time.Now(),
		},
	}

	_, err = userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
