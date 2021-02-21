package graph

import "github.com/SpencerWhitehead7/no-manga/server/graph/repository"

// To run go generate recursively over entire project, use this command: go generate ./...
//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver serves as dependency injection for no-manga's gqlResolvers.
type Resolver struct {
	MangaRepo *repository.Manga
}
