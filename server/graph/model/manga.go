package model

import "time"

// Manga represents a row in the manga table.
type Manga struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	OtherNames  []string   `json:"otherNames"`
	Description string     `json:"description"`
	Demo        string     `json:"demo"`
	StartDate   time.Time  `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
}
