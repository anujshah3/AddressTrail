package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/anujshah3/AddressTrail/internal/handlers"
	"github.com/anujshah3/AddressTrail/internal/middleware"
	"github.com/gorilla/sessions"
)



var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))


func handleIndex(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "web/templates/index.html")
}

func main() {
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