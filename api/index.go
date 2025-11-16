package handler

import (
	"context"
	"net/http"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"sct-backend-service/app/options/config"
	"sct-backend-service/app/options/data"
	"sct-backend-service/app/options/service"
	"sct-backend-service/app/workflow"
	"sct-backend-service/graph"
	"sct-backend-service/graph/generated"
)

var (
	graphqlHandler    *handler.Server
	playgroundHandler http.Handler
	initOnce          sync.Once
	appInstance       *fx.App
)

func initializeHandlers() {
	initOnce.Do(func() {
		// Create resolver
		resolver := &graph.Resolver{}

		// Use fx to extract workflow service
		var workflowService workflow.WorkflowGraphQLService
		var logger *zap.Logger

		// Create fx container with all dependencies (excluding HTTP server for serverless)
		appInstance = fx.New(
			config.ConfigFxOption(""),
			config.LoggerFxOption(),
			data.QueryFxOption(),
			service.ControllerFxOption(),
			service.WorkflowFxOption(),
			// Note: We don't include http.HttpFxOption() for serverless
			fx.Invoke(func(w workflow.WorkflowGraphQLService, l *zap.Logger) {
				workflowService = w
				logger = l
			}),
		)

		// Start the app to trigger Invoke functions (this doesn't start HTTP server)
		ctx := context.Background()
		if err := appInstance.Start(ctx); err != nil {
			// Fallback: create minimal logger
			logger, _ = zap.NewProduction()
			logger.Error("Failed to start fx app", zap.Error(err))
		}

		// Set workflow on resolver
		if workflowService != nil {
			resolver.Workflow = workflowService
		}

		// Create GraphQL handler
		config := generated.Config{
			Resolvers: resolver,
		}
		config.Directives = generated.DirectiveRoot{}

		executableSchema := generated.NewExecutableSchema(config)
		graphqlHandler = handler.NewDefaultServer(executableSchema)
		playgroundHandler = playground.Handler("GraphQL Playground", "/api/graphql")

		_ = logger
	})
}

// Handler is the Vercel serverless function entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	initializeHandlers()

	path := r.URL.Path

	switch path {
	case "/api/playground", "/api/":
		if playgroundHandler != nil {
			playgroundHandler.ServeHTTP(w, r)
		} else {
			http.Error(w, "Playground not available", http.StatusNotFound)
		}
	case "/api/graphql", "/api/query":
		if graphqlHandler != nil {
			graphqlHandler.ServeHTTP(w, r)
		} else {
			http.Error(w, "GraphQL handler not initialized", http.StatusInternalServerError)
		}
	default:
		// Default to GraphQL endpoint
		if graphqlHandler != nil {
			graphqlHandler.ServeHTTP(w, r)
		} else {
			http.Error(w, "GraphQL handler not initialized", http.StatusInternalServerError)
		}
	}
}
