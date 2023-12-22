package mortar

import (
	"github.com/aku-ato/scanner/app/customhandlers"
	"github.com/aku-ato/scanner/app/customhandlers/openapi"
	"go.uber.org/fx"
)

func CustomHandlersFxOptions() fx.Option {
	return fx.Options(
		fx.Provide(openapi.CreateDocumentationHandler),
		customhandlers.FxOption(),
	)
}
