package handlers

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/anujshah3/AddressTrail/config"
	"github.com/anujshah3/AddressTrail/internal/middleware"
	"github.com/anujshah3/AddressTrail/internal/models"
	"github.com/anujshah3/AddressTrail/internal/services"
)

// func GoogleLoginHandler(res http.ResponseWriter, req *http.Request){
// 	session, _ := middleware.GetSession(req, "session")

// 	if middleware.IsAuthenticated(session) {
// 		http.Redirect(res, req, "/dashboard", http.StatusFound)
// 		return
// 	}

// 	googleConfig := config.SetupConfig()
// 	RandomString := os.Getenv("RANDOM_STRING")

// 	url := googleConfig.AuthCodeURL(RandomString)

// 	http.Redirect(res, req, url, http.StatusSeeOther)
// }
func GoogleLoginHandler(c *gin.Context) {
	if middleware.IsAuthenticated(c) {
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}

	googleConfig := config.SetupConfig()
	RandomString := os.Getenv("RANDOM_STRING")

	url := googleConfig.AuthCodeURL(RandomString)

	log.Printf("Redirecting to URL: %s\n", url)

	c.Redirect(http.StatusSeeOther, url)
}


// func GoogleCallBackHandler(res http.ResponseWriter, req *http.Request){
// 	// Access environment variables
// 	RandomString := os.Getenv("RANDOM_STRING")	
	
// 	state := req.URL.Query()["state"][0]
// 	if state != RandomString {
// 		fmt.Fprintf(res, "states don't match")
// 		return
// 	}

// 	code := req.URL.Query()["code"][0]

// 	googleConfig := config.SetupConfig()

// 	token, err := googleConfig.Exchange(context.Background(), code)
// 	if err != nil {
// 		fmt.Fprintln(res, "Code token exchange failed!")
// 	}

// 	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token="+token.AccessToken)
// 	if err != nil {
// 		fmt.Fprintln(res, "User Data Fetch failed!")
// 	}

// 	userDataByte, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Fprintln(res, "JSON Data Parsing failed!")
// 		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
// 		fmt.Println(err.Error())
// 	}

// 	var userData map[string]interface{}

// 	err = json.Unmarshal(userDataByte, &userData)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// session, err := middleware.GetSession(req, "session")
// 	// if err != nil {
// 	// 	fmt.Fprintln(res, err, "Failed to create session")
// 	// 	return
// 	// }
	
// 	gob.Register(userData)

// 	user := &models.User{
// 		ID:    uuid.New().String(),
// 		Name:  userData["name"].(string),
// 		Email: userData["email"].(string),
// 		Addresses: []*models.AddressWithDates{},
// 	}
// 	userID, err := services.AddUser(user)
// 	if err != nil {
// 		fmt.Fprintln(res, "Failed to add user to the database")
// 		return
// 	}

// 	// middleware.SetAuthenticated(session, userID)

// 	// err = session.Save(req, res)
// 	// if err != nil {
// 	// 	fmt.Fprintln(res, err, "Failed to save session")
// 	// 	return
// 	// }
// 	http.Redirect(res, req, "/dashboard", http.StatusFound)
// }

func GoogleCallBackHandler(c *gin.Context) {
	// Access environment variables
	RandomString := os.Getenv("RANDOM_STRING")

	state := c.Request.URL.Query().Get("state")
	if state != RandomString {
		c.String(http.StatusBadRequest, "states don't match")
		return
	}

	code := c.Request.URL.Query().Get("code")

	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Code token exchange failed!")
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.String(http.StatusInternalServerError, "User Data Fetch failed!")
		return
	}

	userDataByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "JSON Data Parsing failed!")
		fmt.Println(err.Error())
		return
	}

	var userData map[string]interface{}

	err = json.Unmarshal(userDataByte, &userData)
	if err != nil {
		fmt.Println("Error:", err)
		c.String(http.StatusInternalServerError, "Failed to parse JSON data")
		return
	}

	gob.Register(userData)

	user := &models.User{
		ID:        uuid.New().String(),
		Name:      userData["name"].(string),
		Email:     userData["email"].(string),
		Addresses: []*models.AddressWithDates{},
	}

	userID, err := services.AddUser(user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to add user to the database")
		return
	}

	middleware.CreateSession(c, userID, user.Name)

	log.Printf("Authenticated Successfully:!\n")

	c.Redirect(http.StatusFound, "/dashboard")
}
