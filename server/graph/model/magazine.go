package model

// Magazine represents a row in the magazine table.
type Magazine struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	OtherNames  []string `json:"otherNames"`
	Description string   `json:"description"`
	Demo        string   `json:"demo"`
}
