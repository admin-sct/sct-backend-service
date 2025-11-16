package entities

// Domain entities
// These represent the core business objects

// User represents a user entity
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt string
}

// ToModel converts entity to GraphQL model
// This will be implemented when models are imported

