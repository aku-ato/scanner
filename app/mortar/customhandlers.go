package mortar

import (
	"com.facehunter/scanner/app/customhandlers"
	"com.facehunter/scanner/app/customhandlers/openapi"
	"go.uber.org/fx"
)

func CustomHandlersFxOptions() fx.Option {
	return fx.Options(
		fx.Provide(openapi.CreateDocumentationHandler),
		customhandlers.FxOption(),
	)
}
