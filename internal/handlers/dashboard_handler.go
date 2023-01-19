package handlers

import (
	"fmt"
	"net/http"
)



func DashboardHandler(res http.ResponseWriter, req *http.Request) {
    session, _ := store.Get(req, "user-session")

    // Check if user is authenticated
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(res, "Forbidden", http.StatusForbidden)
        return
    }

    if userData, ok := session.Values["userData"].(string); ok {
        fmt.Println("User Data:", userData)
    }
    
	http.ServeFile(res, req, "web/templates/dashboard.html")
}
