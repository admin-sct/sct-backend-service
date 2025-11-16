package workflow

import (
	"context"
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"sct-backend-service/app/controllers"
	"sct-backend-service/graph/model"
	"sct-backend-service/types"
)

type WorkflowGraphQLService interface {
	types.GraphQLService
}

type WorkflowGraphQLServiceDeps struct {
	fx.In
	Logger     *zap.Logger
	Controller controllers.GraphQLController
}

type workflowGraphQLServiceDepsImpl struct {
	deps WorkflowGraphQLServiceDeps
}

func CreateWorkflowGraphQLService(deps WorkflowGraphQLServiceDeps) WorkflowGraphQLService {
	return &workflowGraphQLServiceDepsImpl{
		deps: deps,
	}
}

func (impl *workflowGraphQLServiceDepsImpl) SendContactInfo(ctx context.Context, input model.SendContactInfoRequest) (*model.SendContactInfoResponse, error) {
	// Add tracing/observability here
	impl.deps.Logger.Info("SendContactInfo workflow started",
		zap.Int("contact_count", len(input.ContactInfo)),
	)

	// TODO: Add tracing span
	// ctx, span := impl.deps.Tracer.Start(ctx, "SendContactInfo/workflow")
	// defer span.End()

	// Delegate to controller
	result, err := impl.deps.Controller.SendContactInfo(ctx, input)
	if err != nil {
		impl.deps.Logger.Error("SendContactInfo workflow failed",
			zap.Error(err),
		)
		return nil, fmt.Errorf("workflow error: %w", err)
	}

	impl.deps.Logger.Info("SendContactInfo workflow completed",
		zap.Bool("success", result.IsSuccess),
	)

	return result, nil
}
