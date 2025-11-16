package keys

// Application constants

const (
	// Default values
	DefaultPort     = 8080
	DefaultHost     = "0.0.0.0"
	DefaultLogLevel = "info"

	// Paths
	GraphQLPath     = "/query"
	PlaygroundPath  = "/"
	HealthCheckPath = "/health"

	// Timeouts (in seconds)
	DefaultReadTimeout  = 15
	DefaultWriteTimeout = 15
	DefaultIdleTimeout  = 60
)

