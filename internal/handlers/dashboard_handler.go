package handlers

import (
	"fmt"
	"net/http"

	"github.com/anujshah3/AddressTrail/internal/middleware"
)



func DashboardHandler(res http.ResponseWriter, req *http.Request) {
	session, _ := middleware.GetSession(req, "user-session")

	if !middleware.IsAuthenticated(session) {
		http.Error(res, "Forbidden", http.StatusForbidden)
		return
	}

    userInfo := middleware.GetUserInfo(session)
    fmt.Println("User Data:", userInfo)
    
	http.ServeFile(res, req, "web/templates/dashboard.html")
}
