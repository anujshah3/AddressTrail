package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/anujshah3/AddressTrail/config"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request){
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")

	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallBack(res http.ResponseWriter, req *http.Request){

	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
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

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(res, "JSON Data Parsing failed!")
	}

	fmt.Fprintln(res, string(userData))
}