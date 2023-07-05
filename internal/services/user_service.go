package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/anujshah3/AddressTrail/config"
	"github.com/anujshah3/AddressTrail/internal/models"
)

func GetUserDetails(userID string) (models.User, error) {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return models.User{}, err
	}
	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")
	user := models.User{}
	err = userCollection.FindOne(context.Background(), bson.M{"id": userID}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

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
		_, err := userCollection.InsertOne(context.Background(), user, opts)
		if err != nil {
			return "", err
		}
		return user.ID, nil
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

	_, err = userCollection.DeleteOne(context.Background(), bson.M{"id": userID})
	if err != nil {
		return err
	}
	return nil
}

func UpdateCurrentAddressFlag(userID string, endDate time.Time) error {
    client, err := config.GetMongoDBClient()
    if err != nil {
        return err
    }
    defer client.Disconnect(context.Background())
    
    userCollection := config.GetCollection(client, "user")
    
    filter := bson.M{"id": userID, "addresses.current": true}
    update := bson.M{"$set": bson.M{"addresses.$.current": false, "addresses.$.endDate": endDate}}
    
    _, err = userCollection.UpdateOne(context.Background(), filter, update)
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



func GetUserAddresses(userID string) ([]*models.UserAddressesResponse, error) {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	userCollection := config.GetCollection(client, "user")
	filter := bson.M{"id": userID}
	projection := bson.M{"addresses.addressid": 1, "addresses.startdate": 1, "addresses.enddate": 1}

	result := []*models.UserAddressesResponse{}
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

			userAddress := &models.UserAddressesResponse{
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
		"id": userID,
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
