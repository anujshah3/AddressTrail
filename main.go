package main

import (
	"fmt"
	"net/http"

	"github.com/anujshah3/AddressTrail/controller"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/index.html")
}

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/login", controller.GoogleLogin)
	http.HandleFunc("/dashboard", controller.GoogleCallBack)

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}