package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/anujshah3/AddressTrail/internal/middleware"
)


func AddressBookHandler(res http.ResponseWriter, req *http.Request) {
	session, _ := middleware.GetSession(req, "user-session")

	if !middleware.IsAuthenticated(session) {
		http.Error(res, "Forbidden", http.StatusForbidden)
		return
	}

    userInfo := middleware.GetUserInfo(session)
    fmt.Println("User Data:", userInfo["given_name"])
	userName := fmt.Sprint(userInfo["given_name"])
	data := PageData{
		Name: userName,
	}

	tmpl, err := template.ParseFiles("web/templates/address-book.html")
	
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template with the data
	err = tmpl.Execute(res, data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// http.ServeFile(res, req, "web/templates/dashboard.html")
}
