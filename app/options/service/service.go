package service

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"sct-backend-service/app/controllers"
	"sct-backend-service/app/query"
	"sct-backend-service/app/workflow"
)

// ControllerFxOption provides controller dependencies via fx
func ControllerFxOption() fx.Option {
	return fx.Options(
		fx.Provide(NewGraphQLController),
	)
}

// WorkflowFxOption provides workflow dependencies via fx
func WorkflowFxOption() fx.Option {
	return fx.Options(
		fx.Provide(workflow.CreateWorkflowGraphQLService),
	)
}

// NewGraphQLController creates a new GraphQL controller with dependencies
func NewGraphQLController(
	logger *zap.Logger,
	queryBuilder *query.QueryBuilder,
) controllers.GraphQLController {
	deps := controllers.ControllerDeps{
		Logger:       logger,
		QueryBuilder: queryBuilder,
	}

	return controllers.CreateGraphQLController(deps)
}
