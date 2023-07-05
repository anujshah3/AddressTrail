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

// func GetCoordinates(address string) (float64, float64, error) {
// 	apiURL := "https://maps.googleapis.com/maps/api/geocode/json"
// 	apiKey := "API_KEY"

// 	queryParams := url.Values{}
// 	queryParams.Set("address", address)
// 	queryParams.Set("key", apiKey)
// 	requestURL := fmt.Sprintf("%s?%s", apiURL, queryParams.Encode())

// 	response, err := http.Get(requestURL)
// 	if err != nil {
// 		return 0, 0, err
// 	}
// 	defer response.Body.Close()

// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	var data struct {
// 		Results []struct {
// 			Geometry struct {
// 				Location struct {
// 					Lat float64
// 					Lng float64
// 				}
// 			}
// 		}
// 	}
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	if len(data.Results) > 0 {
// 		lat := data.Results[0].Geometry.Location.Lat
// 		lng := data.Results[0].Geometry.Location.Lng
// 		return lat, lng, nil
// 	}

// 	return 0, 0, fmt.Errorf("no results found")
// }
