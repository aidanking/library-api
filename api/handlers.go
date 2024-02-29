package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/aidanking/library-api/storage"
	"github.com/aidanking/library-api/types"
	"github.com/gin-gonic/gin"
)

type authorHandlers struct {
	authorRepository *storage.AuthorRepository
}

func (handlers *authorHandlers) handleCreateAuthor(c *gin.Context) {
	var payload types.Author

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorMessage{ErrorMessage: "bad request data"})
		return
	}

	validateErr := payload.Validate()

	if validateErr != nil {
		c.JSON(http.StatusUnprocessableEntity, types.ErrorMessage{ErrorMessage: validateErr.Error()})
		return
	}

	createdAuthor, createErr := handlers.authorRepository.CreateAuthor(&payload)

	if createErr != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorMessage{ErrorMessage: "error while creating author data"})
		return
	}

	c.JSON(http.StatusCreated, createdAuthor)
}

func (handlers *authorHandlers) handleGetAuthor(c *gin.Context) {

	id, idErr := strconv.Atoi(c.Param("id"))

	if idErr != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorMessage{ErrorMessage: "internal server error"})
		return
	}

	author, authorErr := handlers.authorRepository.FindAuthorById(int64(id))

	if authorErr != nil && authorErr == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, types.ErrorMessage{ErrorMessage: "author not found"})
		return
	} else if authorErr != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorMessage{ErrorMessage: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func (handlers *authorHandlers) handleGetAuthors(c *gin.Context) {

	authors, authorsErr := handlers.authorRepository.FindAllAuthors()

	if authorsErr != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorMessage{ErrorMessage: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func (handlers *authorHandlers) handleUpdateAuthor(c *gin.Context) {

	var payload types.Author

	if bindErr := c.BindJSON(&payload); bindErr != nil {
		c.JSON(http.StatusBadRequest, types.ErrorMessage{ErrorMessage: "bad request data"})
		return
	}

	validateErr := payload.Validate()

	if validateErr != nil {
		c.JSON(http.StatusUnprocessableEntity, types.ErrorMessage{ErrorMessage: validateErr.Error()})
		return
	}

	id, numErr := strconv.Atoi(c.Param("id"))

	if numErr != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorMessage{ErrorMessage: "internal server error"})
		return
	}

	updatedAuthor, authorErr := handlers.authorRepository.UpdateAuthor(int64(id), &payload)

	if authorErr != nil && authorErr == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, types.ErrorMessage{ErrorMessage: "author not found"})
		return
	}

	if authorErr != nil {
		c.JSON(http.StatusNotFound, types.ErrorMessage{ErrorMessage: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, updatedAuthor)
}

func (handlers *authorHandlers) handleDeleteAuthor(c *gin.Context) {
	id, idErr := strconv.Atoi(c.Param("id"))

	if idErr != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorMessage{ErrorMessage: "internal server Error"})
		return
	}

	deletedAuthor, authorErr := handlers.authorRepository.DeleteAuthor(int64(id))

	if authorErr != nil && authorErr == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, types.ErrorMessage{ErrorMessage: "author not found"})
		return
	} else if authorErr != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorMessage{ErrorMessage: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, deletedAuthor)
}
