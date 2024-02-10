package api

type AuthorRequestPayload struct {
	Author AuthorRequestData `json:"author"`
}

type AuthorRequestData struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName"`
	Country    string `json:"country"`
}

type AuthorsPayload struct {
	Authors []AuthorData `json:"authors"`
}

type AuthorPayload struct {
	Author AuthorData `json:"author"`
}

type AuthorData struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName"`
	Country    string `json:"country"`
}

type ErrorMessage struct {
	ErrorMessage string `json:"errorMessage"`
}
