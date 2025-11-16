package options

import (
	"go.uber.org/fx"

	"sct-backend-service/app/options/config"
	"sct-backend-service/app/options/data"
	"sct-backend-service/app/options/http"
	"sct-backend-service/app/options/service"
)

// CreateApplication creates the fx application with all dependencies
func CreateApplication(configFilePath string) *fx.App {
	return fx.New(
		config.ConfigFxOption(configFilePath),
		config.LoggerFxOption(),
		data.QueryFxOption(),
		service.ControllerFxOption(),
		service.WorkflowFxOption(),
		http.HttpFxOption(),
	)
}
