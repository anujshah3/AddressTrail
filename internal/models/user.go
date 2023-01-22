package models

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Addresses []*AddressWithDates
}

type AddressWithDates struct {
	Address   *Address
	StartDate time.Time
	EndDate   time.Time
}
