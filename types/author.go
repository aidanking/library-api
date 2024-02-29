package types

import (
	"fmt"
)

type Author struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName"`
	Country    string `json:"country"`
}

func (author *Author) Validate() error {
	if author.FirstName == "" {
		return fmt.Errorf("first name is required")
	}

	maxLength := 255

	if len(author.FirstName) > maxLength {
		return fmt.Errorf("first name should be less than %d", maxLength)
	}

	if author.MiddleName != "" && len(author.MiddleName) > maxLength {
		return fmt.Errorf("middle name should be less than or equal to %d", maxLength)
	}

	if author.LastName == "" {
		return fmt.Errorf("last name is required")

	}

	if len(author.LastName) > maxLength {
		return fmt.Errorf("last name should be less than or equal to %d", maxLength)
	}

	if author.Country == "" {
		return fmt.Errorf("author country is required")
	}

	if len(author.Country) > maxLength {
		return fmt.Errorf("country length should be less than or equal to %d", maxLength)
	}

	return nil
}

type ErrorMessage struct {
	ErrorMessage string `json:"errorMessage"`
}
