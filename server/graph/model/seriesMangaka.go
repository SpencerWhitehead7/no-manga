package model

// SeriesMangaka represents a row in the mangaka table, with mangaka's job on that series.
type SeriesMangaka struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	OtherNames  []string `json:"otherNames"`
	Description string   `json:"description"`
	Job         string   `json:"job"`
}
