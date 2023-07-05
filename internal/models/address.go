package models

type Address struct {
	Street     string
	Unit       string
	City       string
	State      string
	PostalCode string
	Country    string
	Latitude   float64
	Longitude  float64
}
