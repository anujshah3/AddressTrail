package models

import (
	"time"
)

type Usercompanydata struct {
	UserID       string
	CompanyData  []CompanyDataEntry
}

type CompanyDataEntry struct {
	CompanyID     string
	LastRequest   time.Time
	CallbackURL   string
}

type AddCompanyUserDataPayload struct {
	UserID       string
	CompanyID    string
	CallbackURL  string
}
