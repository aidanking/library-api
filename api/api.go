package api

import "github.com/gin-gonic/gin"

var authorsData = []AuthorData{{Id: 1, FirstName: "Stephen", LastName: "King", Country: "US"}, {Id: 2, FirstName: "George", MiddleName: "R.R.", LastName: "Martin"}, {Id: 3, FirstName: "Sarah", MiddleName: "J.", LastName: "Mass", Country: "US"}, {Id: 4, FirstName: "Brandon", LastName: "Sanderson", Country: "US"}, {Id: 5, FirstName: "Emily", LastName: "Henry", Country: "US"}}

func ServerStart(listenAddr string) {

	r := gin.Default()

	api := r.Group("/api")

	api.POST("/authors", handleCreateAuthor)
	api.GET("/authors", handleGetAuthors)
	api.GET("/authors/:id", handleGetAuthor)
	api.PUT("/authors/:id", handleUpdateAuthor)
	api.DELETE("/authors/:id", handleDeleteAuthor)

	r.Run(listenAddr)

}
