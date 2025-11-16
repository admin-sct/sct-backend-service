package controllers

import (
	"go.uber.org/zap"

	"sct-backend-service/app/query"
)

// ControllerDeps holds shared dependencies for controllers
type ControllerDeps struct {
	Logger       *zap.Logger
	QueryBuilder *query.QueryBuilder
}
