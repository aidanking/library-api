package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleCreateAuthor(c *gin.Context) {
	var payload AuthorRequestPayload

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{ErrorMessage: "bad request data"})
		return
	}

	createdAuthor := createAuthor(payload.Author)

	c.JSON(http.StatusCreated, AuthorPayload{Author: createdAuthor})
}

func handleGetAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, AuthorsPayload{Authors: authorsData})
}

func handleGetAuthor(c *gin.Context) {

	id, idErr := strconv.Atoi(c.Param("id"))

	if idErr != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage{ErrorMessage: "internal server error"})
		return
	}

	author, _, authorErr := findAuthorById(id)

	if authorErr != nil {
		c.JSON(http.StatusNotFound, ErrorMessage{ErrorMessage: authorErr.Error()})
		return
	}

	c.JSON(http.StatusOK, AuthorPayload{Author: author})
}

func handleUpdateAuthor(c *gin.Context) {

	var payload AuthorRequestPayload

	if bindErr := c.BindJSON(&payload); bindErr != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{ErrorMessage: "bad request data"})
		return
	}

	id, numErr := strconv.Atoi(c.Param("id"))

	if numErr != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage{ErrorMessage: "internal server error"})
		return
	}

	updatedAuthor, authorErr := updateAuthor(id, payload.Author)

	if authorErr != nil {
		c.JSON(http.StatusNotFound, ErrorMessage{ErrorMessage: authorErr.Error()})
		return
	}

	c.JSON(http.StatusOK, AuthorPayload{Author: updatedAuthor})
}

func handleDeleteAuthor(c *gin.Context) {
	id, idErr := strconv.Atoi(c.Param("id"))

	if idErr != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage{ErrorMessage: "internal server Error"})
		return
	}

	deletedAuthor, authorErr := deleteAuthor(id)

	if authorErr != nil {
		c.JSON(http.StatusNotFound, ErrorMessage{ErrorMessage: authorErr.Error()})
		return
	}

	c.JSON(http.StatusOK, AuthorPayload{Author: deletedAuthor})
}

func createAuthor(requestAuthor AuthorRequestData) AuthorData {

	newData := AuthorData{Id: 0, FirstName: requestAuthor.FirstName, MiddleName: requestAuthor.MiddleName, LastName: requestAuthor.LastName, Country: requestAuthor.Country}
	authorsData = append(authorsData, newData)

	return newData
}

func updateAuthor(id int, requestAuthor AuthorRequestData) (AuthorData, error) {
	_, authorIndex, authorErr := findAuthorById(id)

	if authorErr != nil {
		return AuthorData{}, authorErr
	}

	newData := AuthorData{Id: id, FirstName: requestAuthor.FirstName, MiddleName: requestAuthor.MiddleName, LastName: requestAuthor.LastName, Country: requestAuthor.Country}
	authorsData[authorIndex] = newData

	return newData, nil
}

func findAuthorById(id int) (AuthorData, int, error) {

	for authorIndex, author := range authorsData {
		if author.Id == id {
			return author, authorIndex, nil
		}
	}

	return AuthorData{}, -1, errors.New("author not found")
}

func deleteAuthor(id int) (AuthorData, error) {

	author, authorIndex, authorErr := findAuthorById(id)

	if authorErr != nil {
		return AuthorData{}, authorErr
	}

	authorsData = append(authorsData[:authorIndex], authorsData[authorIndex+1:]...)

	return author, nil
}
