// Package repos provides a common interface for repository implementations.
package repos

import "io"

// Basic Repository representation.
type Repository struct {
	Name string `json:"name"`
}

// Basic package representation.
type Package struct {
	Name      string `json:"name"`
	Timestamp string `json:"lastUpdated"`
	Size      int64  `json:"size"`
}

type RepositoryStore interface {
	// Lists all repositories in store.
	ListRepositories() []Repository

	Initialize(name string) error
	Get(name string) AnyRepository
}

type AnyRepository interface {
	Add(filename string, f io.Reader)
	Update() error
	ListPackages() []Package
	IsValid() bool
}
