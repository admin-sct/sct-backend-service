package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"sct-backend-service/app/options/config"
	"sct-backend-service/app/workflow"
	"sct-backend-service/graph"
	"sct-backend-service/internal/server"
)

// HTTPServer holds the HTTP server instance
type HTTPServer struct {
	server *server.Server
}

// NewHTTPServer creates a new HTTP server with all dependencies
func NewHTTPServer(
	lc fx.Lifecycle,
	cfg *config.Config,
	logger *zap.Logger,
	workflowService workflow.WorkflowGraphQLService,
) (*HTTPServer, error) {
	// Create resolver
	resolver := &graph.Resolver{
		Workflow: workflowService,
	}

	// Build server
	srv, err := server.NewServerBuilder().
		WithPort(cfg.Server.Port).
		WithHost(cfg.Server.Host).
		WithReadTimeout(15 * time.Second).
		WithWriteTimeout(15 * time.Second).
		WithIdleTimeout(60 * time.Second).
		WithPlayground(true).
		WithPlaygroundPath("/").
		WithGraphQLPath("/query").
		WithResolvers(resolver).
		Build()

	if err != nil {
		return nil, fmt.Errorf("failed to build server: %w", err)
	}

	httpServer := &HTTPServer{
		server: srv,
	}

	// Register lifecycle hooks
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.Start(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Server failed to start: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			return srv.Shutdown(shutdownCtx)
		},
	})

	return httpServer, nil
}

// HttpFxOption provides HTTP server dependencies via fx
func HttpFxOption() fx.Option {
	return fx.Options(
		fx.Provide(NewHTTPServer),
		fx.Invoke(func(*HTTPServer) {}), // Invoke to start the server
	)
}
