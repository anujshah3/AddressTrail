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


func AddCompanyData(data *models.AddCompanyUserDataPayload) (string, error) {
	client, err := config.GetMongoDBClient()
	if err != nil {
		return "", err
	}
	defer client.Disconnect(context.Background())

	companyDataCollection := config.GetCollection(client, "Usercompanydata")
	existingData := models.Usercompanydata{}
	err = companyDataCollection.FindOne(context.Background(), bson.M{"userID": data.UserID}).Decode(&existingData)
	if err == mongo.ErrNoDocuments {
		opts := options.InsertOne()
		_, err := companyDataCollection.InsertOne(context.Background(), &models.Usercompanydata{
			UserID:      data.UserID,
			CompanyData: []models.CompanyDataEntry{{CompanyID: data.CompanyID, LastRequest: time.Now(), CallbackURL: data.CallbackURL}},
		}, opts)
		if err != nil {
			return "", err
		}
		return data.UserID, nil
	}

	for i := range existingData.CompanyData {
		if existingData.CompanyData[i].CompanyID == data.CompanyID {
			if existingData.CompanyData[i].CallbackURL != data.CallbackURL {
				existingData.CompanyData[i].LastRequest = time.Now()
				existingData.CompanyData[i].CallbackURL = data.CallbackURL
				update := bson.M{
					"$set": bson.M{
						"companyData": existingData.CompanyData,
					},
				}
				_, err := companyDataCollection.UpdateOne(context.Background(), bson.M{"userID": data.UserID}, update)
				if err != nil {
					return "", err
				}
			}
			return existingData.UserID, nil
		}
	}

	existingData.CompanyData = append(existingData.CompanyData, models.CompanyDataEntry{
		CompanyID:    data.CompanyID,
		LastRequest:  time.Now(),
		CallbackURL:  data.CallbackURL,
	})

	update := bson.M{
		"$set": bson.M{
			"companyData": existingData.CompanyData,
			"userID":      existingData.UserID,
		},
	}
	_, err = companyDataCollection.UpdateOne(context.Background(), bson.M{"userID": data.UserID}, update)
	if err != nil {
		return "", err
	}

	return existingData.UserID, nil
}
