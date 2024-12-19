package store

// TODO: get rid of database dependency
import "zeta/internal/cache/database"

type Store interface {
	// Core Operations
	UpdateOne(path string) error
	UpdateAll() error
	Recompute() error

	// Queries
	GetAll() ([]string, error)
	GetParents(path string) ([]string, error)
	GetChildren(path string) ([]string, error)

	// TODO: Change to store specific signature
	GetGraph() ([]database.FileRecord, []database.LinkRecord, error)
}
