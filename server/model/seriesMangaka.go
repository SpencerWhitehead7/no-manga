package model

// Mangaka + job field (job refers to their role on a specific series)
// so it only makes sense in the context of a Manga
type SeriesMangaka struct {
	Mangaka
	Job string `json:"job"`
}
