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
	Current     bool
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserAddressesResponse struct {
	AddressID string
	Current     bool
	Street     string
	Unit       string
	City       string
	State      string
	PostalCode string
	Country    string
	StartDate time.Time
	EndDate   time.Time
}