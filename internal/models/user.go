package models

import "time"

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
