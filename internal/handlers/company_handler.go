package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyData struct {
	CompanyName string
}


func CompanyDashboardHandler(c *gin.Context) {
	companyName := c.Query("name")

	data := CompanyData{
		CompanyName: companyName,
	}

	tmpl, err := template.ParseFiles("web/templates/company.html")
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
