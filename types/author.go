package types

type Author struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName"`
	Country    string `json:"country"`
}

type ErrorMessage struct {
	ErrorMessage string `json:"errorMessage"`
}
