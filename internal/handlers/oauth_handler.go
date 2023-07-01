package handlers

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/uuid"

	"github.com/anujshah3/AddressTrail/config"
	"github.com/anujshah3/AddressTrail/internal/middleware"
	"github.com/anujshah3/AddressTrail/internal/models"
	"github.com/anujshah3/AddressTrail/internal/services"
	"github.com/gorilla/sessions"
)


var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

func GoogleLoginHandler(res http.ResponseWriter, req *http.Request){
	session, _ := middleware.GetSession(req, "user-session")

	if middleware.IsAuthenticated(session) {
		http.Redirect(res, req, "/dashboard", http.StatusFound)
		return
	}

	googleConfig := config.SetupConfig()
	RandomString := os.Getenv("RANDOM_STRING")	
		
	url := googleConfig.AuthCodeURL(RandomString)

	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallBackHandler(res http.ResponseWriter, req *http.Request){
	// Access environment variables
	RandomString := os.Getenv("RANDOM_STRING")	
	
	state := req.URL.Query()["state"][0]
	if state != RandomString {
		fmt.Fprintf(res, "states don't match")
		return
	}

	code := req.URL.Query()["code"][0]

	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintln(res, "Code token exchange failed!")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token="+token.AccessToken)
	if err != nil {
		fmt.Fprintln(res, "User Data Fetch failed!")
	}

	userDataByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(res, "JSON Data Parsing failed!")
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		fmt.Println(err.Error())
	}

	var userData map[string]interface{}

	err = json.Unmarshal(userDataByte, &userData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	session, err := middleware.GetSession(req, "user-session")
	if err != nil {
		fmt.Fprintln(res, err, "Failed to create session")
		return
	}
	
	gob.Register(userData)

	user := &models.User{
		ID:    uuid.New().String(),
		Name:  userData["name"].(string),
		Email: userData["email"].(string),
		Addresses: []*models.AddressWithDates{},
	}
	userID, err := services.AddUser(user)
	if err != nil {
		fmt.Fprintln(res, "Failed to add user to the database")
		return
	}

	userData["userID"] = userID
	middleware.SetAuthenticated(session, userData)

	fmt.Println(userData)

	err = session.Save(req, res)
	if err != nil {
		fmt.Fprintln(res, err, "Failed to save session")
		return
	}
	http.Redirect(res, req, "/dashboard", http.StatusFound)
}