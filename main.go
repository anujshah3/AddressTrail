package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/anujshah3/AddressTrail/internal/handlers"
	"github.com/anujshah3/AddressTrail/internal/middleware"
	"github.com/anujshah3/AddressTrail/internal/models"
	"github.com/anujshah3/AddressTrail/internal/services"
	"github.com/gorilla/sessions"
)


var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))


func handleIndex(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "web/templates/index.html")
}

func getRandomDate() time.Time {
	min := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	max := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	delta := max.Sub(min).Seconds()

	randomSeconds := rand.Float64() * delta
	randomDuration := time.Duration(randomSeconds) * time.Second

	return min.Add(randomDuration)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create and add 3 users with 3 addresses each
	id := uuid.New()

	user := &models.User{
		ID:    fmt.Sprintf(id.String()),
		Name:  fmt.Sprintf("Anuj Shah"),
		Email: fmt.Sprintf("anujshah031198@gmail.com"),
	}

	Add_id := uuid.New()

	address := &models.Address{
		ID:         fmt.Sprintf(Add_id.String()),
		Street:     fmt.Sprintf("526 Palace Ct"),
		City:       fmt.Sprintf("Schaumburg"),
		State:      fmt.Sprintf("IL"),
		PostalCode: fmt.Sprintf("60194"),
		Country:    fmt.Sprintf("USA"),
	}

	err := services.InsertAddress(address)
	if err != nil {
		fmt.Printf("Error adding address %s: %s\n", address.ID, err.Error())
	}
			
	addressWithDates := &models.AddressWithDates{
		AddressID:   address.ID,
		StartDate: getRandomDate(),
		EndDate:   getRandomDate(),
	}

	user.Addresses = append(user.Addresses, addressWithDates)

	Add_id = uuid.New()

	address = &models.Address{
		ID:         fmt.Sprintf(Add_id.String()),
		Street:     fmt.Sprintf("718 S Carpenter St"),
		City:       fmt.Sprintf("Chicago"),
		State:      fmt.Sprintf("IL"),
		PostalCode: fmt.Sprintf("60607"),
		Country:    fmt.Sprintf("USA"),
	}

	err = services.InsertAddress(address)
	if err != nil {
		fmt.Printf("Error adding address %s: %s\n", address.ID, err.Error())
	}
			
	addressWithDates = &models.AddressWithDates{
		AddressID:   address.ID,
		StartDate: getRandomDate(),
		EndDate:   getRandomDate(),
	}

	user.Addresses = append(user.Addresses, addressWithDates)

	err = services.AddUser(user)
	if err != nil {
		fmt.Printf("Error adding user %s: %s\n", user.ID, err.Error())
	}

	id = uuid.New()

	user = &models.User{
		ID:    fmt.Sprintf(id.String()),
		Name:  fmt.Sprintf("Akshat Shah"),
		Email: fmt.Sprintf("akshatshah@gmail.com"),
	}

	Add_id = uuid.New()

	address = &models.Address{
		ID:         fmt.Sprintf(Add_id.String()),
		Street:     fmt.Sprintf("526 Palace Ct"),
		City:       fmt.Sprintf("Schaumburg"),
		State:      fmt.Sprintf("IL"),
		PostalCode: fmt.Sprintf("60194"),
		Country:    fmt.Sprintf("USA"),
	}

	err = services.InsertAddress(address)
	if err != nil {
		fmt.Printf("Error adding address %s: %s\n", address.ID, err.Error())
	}
			
	addressWithDates = &models.AddressWithDates{
		AddressID:   address.ID,
		StartDate: getRandomDate(),
		EndDate:   getRandomDate(),
	}

	user.Addresses = append(user.Addresses, addressWithDates)

	Add_id = uuid.New()

	address = &models.Address{
		ID:         fmt.Sprintf(Add_id.String()),
		Street:     fmt.Sprintf("718 S Carpenter St"),
		City:       fmt.Sprintf("Chicago"),
		State:      fmt.Sprintf("IL"),
		PostalCode: fmt.Sprintf("60607"),
		Country:    fmt.Sprintf("USA"),
	}

	err = services.InsertAddress(address)
	if err != nil {
		fmt.Printf("Error adding address %s: %s\n", address.ID, err.Error())
	}
			
	addressWithDates = &models.AddressWithDates{
		AddressID:   address.ID,
		StartDate: getRandomDate(),
		EndDate:   getRandomDate(),
	}

	user.Addresses = append(user.Addresses, addressWithDates)

	err = services.AddUser(user)
	if err != nil {
		fmt.Printf("Error adding user %s: %s\n", user.ID, err.Error())
	}

	http.HandleFunc("/", handleIndex)

	http.HandleFunc("/login", handlers.GoogleLoginHandler)
	http.HandleFunc("/auth/google/callback", handlers.GoogleCallBackHandler)
	http.HandleFunc("/dashboard", middleware.AuthMiddleware(handlers.DashboardHandler))
	http.HandleFunc("/address-book", middleware.AuthMiddleware(handlers.AddressBookHandler))

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}