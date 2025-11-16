package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"sct-backend-service/graph"
	"sct-backend-service/graph/generated"
)

// Server represents the GraphQL server
type Server struct {
	httpServer *http.Server
	handler    *handler.Server
	config     *Config
}

// Config holds server configuration
type Config struct {
	Port            int
	Host            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	PlaygroundEnabled bool
	PlaygroundPath  string
	GraphQLPath     string
	Resolvers       *graph.Resolver
}

// ServerBuilder implements the builder pattern for server configuration
type ServerBuilder struct {
	config *Config
}

// NewServerBuilder creates a new server builder with default values
func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{
		config: &Config{
			Port:              8080,
			Host:              "0.0.0.0",
			ReadTimeout:       15 * time.Second,
			WriteTimeout:      15 * time.Second,
			IdleTimeout:       60 * time.Second,
			PlaygroundEnabled: true,
			PlaygroundPath:    "/",
			GraphQLPath:       "/query",
		},
	}
}

// WithPort sets the server port
func (b *ServerBuilder) WithPort(port int) *ServerBuilder {
	b.config.Port = port
	return b
}

// WithHost sets the server host
func (b *ServerBuilder) WithHost(host string) *ServerBuilder {
	b.config.Host = host
	return b
}

// WithReadTimeout sets the read timeout
func (b *ServerBuilder) WithReadTimeout(timeout time.Duration) *ServerBuilder {
	b.config.ReadTimeout = timeout
	return b
}

// WithWriteTimeout sets the write timeout
func (b *ServerBuilder) WithWriteTimeout(timeout time.Duration) *ServerBuilder {
	b.config.WriteTimeout = timeout
	return b
}

// WithIdleTimeout sets the idle timeout
func (b *ServerBuilder) WithIdleTimeout(timeout time.Duration) *ServerBuilder {
	b.config.IdleTimeout = timeout
	return b
}

// WithPlayground enables/disables the GraphQL playground
func (b *ServerBuilder) WithPlayground(enabled bool) *ServerBuilder {
	b.config.PlaygroundEnabled = enabled
	return b
}

// WithPlaygroundPath sets the playground path
func (b *ServerBuilder) WithPlaygroundPath(path string) *ServerBuilder {
	b.config.PlaygroundPath = path
	return b
}

// WithGraphQLPath sets the GraphQL endpoint path
func (b *ServerBuilder) WithGraphQLPath(path string) *ServerBuilder {
	b.config.GraphQLPath = path
	return b
}

// WithResolvers sets the GraphQL resolvers
func (b *ServerBuilder) WithResolvers(resolvers *graph.Resolver) *ServerBuilder {
	b.config.Resolvers = resolvers
	return b
}

// Build creates and configures the server
func (b *ServerBuilder) Build() (*Server, error) {
	if b.config.Resolvers == nil {
		return nil, fmt.Errorf("resolvers are required")
	}

	// Create GraphQL handler
	config := generated.Config{
		Resolvers: b.config.Resolvers,
	}

	// Add any custom directives here
	config.Directives = generated.DirectiveRoot{}

	// Create the executable schema
	executableSchema := generated.NewExecutableSchema(config)

	// Create GraphQL handler
	h := handler.NewDefaultServer(executableSchema)

	// Optional: customize error handling
	// You can set a custom error presenter here if needed

	// Create HTTP handler with routes
	mux := http.NewServeMux()

	// Add GraphQL endpoint
	mux.Handle(b.config.GraphQLPath, h)

	// Add playground if enabled
	if b.config.PlaygroundEnabled {
		mux.Handle(b.config.PlaygroundPath, playground.Handler("GraphQL Playground", b.config.GraphQLPath))
	}

	// Create HTTP server
	addr := fmt.Sprintf("%s:%d", b.config.Host, b.config.Port)
	httpServer := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  b.config.ReadTimeout,
		WriteTimeout: b.config.WriteTimeout,
		IdleTimeout:  b.config.IdleTimeout,
	}

	return &Server{
		httpServer: httpServer,
		handler:    h,
		config:     b.config,
	}, nil
}

// Start starts the server
func (s *Server) Start() error {
	addr := s.httpServer.Addr
	log.Printf("🚀 Server starting on http://%s", addr)
	log.Printf("📊 GraphQL endpoint: http://%s%s", addr, s.config.GraphQLPath)
	if s.config.PlaygroundEnabled {
		log.Printf("🎮 Playground available at: http://%s%s", addr, s.config.PlaygroundPath)
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("🛑 Shutting down server...")
	return s.httpServer.Shutdown(ctx)
}

