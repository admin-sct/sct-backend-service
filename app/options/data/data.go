package data

import (
	"go.uber.org/fx"

	"sct-backend-service/app/query"
)

// QueryFxOption provides query builder dependencies via fx
func QueryFxOption() fx.Option {
	return fx.Options(
		fx.Provide(query.NewQueryBuilder),
	)
}
