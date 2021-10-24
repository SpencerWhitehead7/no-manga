package model

// Mangaka represents a row in the mangaka table.
type Mangaka struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	OtherNames  []string `json:"otherNames"`
	Description string   `json:"description"`
}
