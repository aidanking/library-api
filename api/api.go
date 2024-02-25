package api

import (
	"database/sql"

	"github.com/aidanking/library-api/storage"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	ListenAddr string
	DB         *sql.DB
}

func (apiServer *ApiServer) Start() {

	authorRepository := storage.AuthorRepository{
		DB: apiServer.DB,
	}

	authorHandlers := authorHandlers{
		authorRepository: &authorRepository,
	}

	r := gin.Default()

	api := r.Group("/api")

	api.POST("/authors", authorHandlers.handleCreateAuthor)
	api.GET("/authors", authorHandlers.handleGetAuthors)
	api.GET("/authors/:id", authorHandlers.handleGetAuthor)
	api.PUT("/authors/:id", authorHandlers.handleUpdateAuthor)
	api.DELETE("/authors/:id", authorHandlers.handleDeleteAuthor)

	r.Run(apiServer.ListenAddr)
}
