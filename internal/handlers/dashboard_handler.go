package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageData struct {
	Name string
}

// func DashboardHandler(res http.ResponseWriter, req *http.Request) {
// 	session, _ := middleware.GetSession(req, "session")

// 	if !middleware.IsAuthenticated(session) {
// 		http.Error(res, "Forbidden", http.StatusForbidden)
// 		return
// 	}

//     userID := middleware.GetUserID(session)
// 	fmt.Println(userID)
// 	userName := "Name"
// 	data := PageData{
// 		Name: userName,
// 	}

// 	tmpl, err := template.ParseFiles("web/templates/dashboard.html")
	
// 	if err != nil {
// 		http.Error(res, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = tmpl.Execute(res, data)
// 	if err != nil {
// 		http.Error(res, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

func DashboardHandler(c *gin.Context) {
	// session, _ := middleware.GetSession(c.Request, "session")

	// if !middleware.IsAuthenticated(session) {
	// 	c.AbortWithStatus(http.StatusForbidden)
	// 	return
	// }

	// userID := middleware.GetUserID(session)
	// fmt.Println(userID)
	userName := "Name"
	data := PageData{
		Name: userName,
	}

	tmpl, err := template.ParseFiles("web/templates/dashboard.html")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
