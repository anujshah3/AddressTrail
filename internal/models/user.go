package models

import (
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Addresses []*AddressWithDates
}

type AddressWithDates struct {
	AddressID   string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserAddressesResponse struct {
	AddressID string
	Street     string
	Unit       string
	City       string
	State      string
	PostalCode string
	Country    string
	StartDate time.Time
	EndDate   time.Time
}